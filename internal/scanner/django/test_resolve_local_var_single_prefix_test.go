//ff:func feature=scan type=test control=sequence topic=django
//ff:what resolveLocalVarIncludes — include(localVar) 단일 prefix 전개 검증
package django

import "testing"

func TestResolveLocalVarIncludesSinglePrefix(t *testing.T) {
	got := expandFile(t, `api_urls = [path("x/", v)]
urlpatterns = [path("api/v1/", include(api_urls))]
`)
	if got["/api/v1/x/"] != "v" {
		t.Fatalf("expected /api/v1/x/ -> v, got %v", got)
	}
}
