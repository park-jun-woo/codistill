//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what resolveLocalVarIncludes — 같은 변수 v1/v2/v3 3중 include가 독립 prefix로 전개되는지 검증
package django

import "testing"

func TestResolveLocalVarIncludesTriplePrefix(t *testing.T) {
	got := expandFile(t, `api_urls = [path("checks/", views.checks)]
urlpatterns = [
    path("api/v1/", include(api_urls)),
    path("api/v2/", include(api_urls)),
    path("api/v3/", include(api_urls)),
]
`)
	for _, want := range []string{"/api/v1/checks/", "/api/v2/checks/", "/api/v3/checks/"} {
		if got[want] != "views.checks" {
			t.Errorf("missing %s -> views.checks; got %v", want, got)
		}
	}
}
