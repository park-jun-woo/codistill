//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what resolveLocalVarIncludes — 파라미터 prefix(ping/<uuid>) 하위 로컬변수 라우트 전개 검증
package django

import "testing"

func TestResolveLocalVarIncludesNestedParam(t *testing.T) {
	got := expandFile(t, `uuid_urls = [
    path("fail", views.ping, {"action": "fail"}),
    path("start", views.ping, {"action": "start"}),
    path("log", views.ping, {"action": "log"}),
]
urlpatterns = [path("ping/<uuid:code>/", include(uuid_urls))]
`)
	for _, want := range []string{"/ping/<uuid:code>/fail", "/ping/<uuid:code>/start", "/ping/<uuid:code>/log"} {
		if got[want] != "views.ping" {
			t.Errorf("missing %s; got %v", want, got)
		}
	}
}
