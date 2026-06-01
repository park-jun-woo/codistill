//ff:func feature=scan type=test control=sequence topic=django
//ff:what 괄호/연결(binary_operator) RHS에서 list 피연산자의 path()만 수집하고 변수/static()는 무시하는지 검증한다
package django

import "testing"

func TestCollectFromURLPatternsRHSParen(t *testing.T) {
	src := `urlpatterns = (
    [ re_path(r"^api/", include("x.urls")), path("h/", v) ]
    + plugin_registry.urls + static("media")
)
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	nodes := findAllByType(root, "assignment")
	if len(nodes) == 0 {
		t.Fatalf("no assignment node")
	}
	got := collectFromURLPatternsRHS(nodes[0], []byte(src))
	// Only the two list entries are collected; the variable (plugin_registry.urls)
	// and the static(...) call operand are ignored.
	if len(got) != 2 {
		t.Fatalf("got %d entries, want 2: %+v", len(got), got)
	}
	if !got[0].isInclude || got[0].includeModule != "x.urls" {
		t.Errorf("entry 0: want include x.urls, got %+v", got[0])
	}
	if got[1].isInclude || got[1].viewName != "v" {
		t.Errorf("entry 1: want view v, got %+v", got[1])
	}
}
