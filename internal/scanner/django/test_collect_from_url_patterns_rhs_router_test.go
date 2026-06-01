//ff:func feature=scan type=test control=sequence topic=django
//ff:what urlpatterns = router.urls 형태가 라우터 참조 entry로 수집되는지 검증한다
package django

import "testing"

func TestCollectFromURLPatternsRHS_RouterURLs(t *testing.T) {
	src := "urlpatterns = router.urls\n"
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	nodes := findAllByType(root, "assignment")
	if len(nodes) == 0 {
		t.Fatal("no assignment node")
	}
	got := collectFromURLPatternsRHS(nodes[0], []byte(src))
	if len(got) != 1 {
		t.Fatalf("got %d entries, want 1", len(got))
	}
	if got[0].includeRouterVar != "router" || !got[0].isInclude {
		t.Errorf("expected router-ref entry, got %+v", got[0])
	}
}
