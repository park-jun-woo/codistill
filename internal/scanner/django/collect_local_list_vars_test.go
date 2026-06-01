//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectLocalListVars — 로컬 리스트변수 대입을 변수명→entry로 인덱싱하고 urlpatterns는 제외하는지 검증
package django

import "testing"

func TestCollectLocalListVars(t *testing.T) {
	fi := newTestFileInfo(t, `api_urls = [path("checks/", views.checks), path("status/", views.status)]
not_a_list = 1
urlpatterns = [path("api/v1/", include(api_urls))]
`)
	index := collectLocalListVars(fi)

	if got := len(index["api_urls"]); got != 2 {
		t.Errorf("expected 2 entries for api_urls, got %d", got)
	}
	if index["api_urls"][0].pattern != "checks/" {
		t.Errorf("unexpected first entry: %+v", index["api_urls"][0])
	}
	if _, ok := index["not_a_list"]; ok {
		t.Error("scalar assignment should not be indexed")
	}
	if _, ok := index["urlpatterns"]; ok {
		t.Error("urlpatterns assignment must be excluded from local list-var index")
	}
}
