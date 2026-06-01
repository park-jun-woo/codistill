//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what 괄호 RHS의 re_path 문자열 include가 3단 연쇄로 전개되고 ^api/ prefix가 보존·정규화되는지 검증한다
package django

import "testing"

func TestStringModuleIncludeChain(t *testing.T) {
	config := mkFile(t, "config/urls.py", "config.urls", `
urlpatterns = (
    [ re_path(r"^api/", include("baserow.api.urls", namespace="api")) ]
    + plugin_registry.urls + static("media")
)
`)
	api := mkFile(t, "baserow/api/urls.py", "baserow.api.urls", `
urlpatterns = [
    path("user/", include("baserow.api.user.urls", namespace="user")),
]
`)
	user := mkFile(t, "baserow/api/user/urls.py", "baserow.api.user.urls", `
urlpatterns = [
    re_path(r"^token-auth/$", obtain_jwt_token),
]
`)
	byModule := collectURLs([]fileInfo{config, api, user})

	// config.urls is the sole root; the two included modules are not roots.
	roots := findRootURLModules(byModule)
	if len(roots) != 1 || roots[0] != "config.urls" {
		t.Fatalf("want root [config.urls], got %v", roots)
	}

	var openapiPaths []string
	for _, r := range roots {
		for _, e := range expandURLModule(r, "", byModule, map[string]bool{}) {
			openapiPaths = append(openapiPaths, djangoURLToOpenAPI(e.pattern))
		}
	}
	// 3-level chain config -> api -> user, with the ^api/ prefix preserved and
	// regex anchors normalized away at every position.
	found := false
	for _, p := range openapiPaths {
		if p == "/api/user/token-auth/" {
			found = true
		}
	}
	if !found {
		t.Errorf("want /api/user/token-auth/ in expansion, got %v", openapiPaths)
	}
}
