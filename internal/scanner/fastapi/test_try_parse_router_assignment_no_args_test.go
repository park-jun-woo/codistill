//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryParseRouterAssignment_NoArgs 테스트
package fastapi

import "testing"

func TestTryParseRouterAssignment_NoArgs(t *testing.T) {

	src := []byte("app = FastAPI()\n")
	root, _ := parsePython(src)
	a := findAllByType(root, "assignment")[0]
	ri := tryParseRouterAssignment(a, src, nil)
	if ri == nil || !ri.isFastAPI || ri.varName != "app" {
		t.Fatalf("FastAPI assignment: got %v", ri)
	}
	if ri.prefix != "" {
		t.Fatalf("expected empty prefix, got %q", ri.prefix)
	}
}
