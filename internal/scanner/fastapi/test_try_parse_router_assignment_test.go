//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryParseRouterAssignment 테스트
package fastapi

import "testing"

func TestTryParseRouterAssignment(t *testing.T) {
	src := []byte("app = FastAPI()\nrouter = APIRouter(prefix='/api')\nother = SomeClass()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	assigns := findAllByType(root, "assignment")
	found := 0
	for _, a := range assigns {
		ri := tryParseRouterAssignment(a, src, nil)
		if ri != nil {
			found++
		}
	}
	if found < 2 {
		t.Fatalf("expected at least 2 router assignments, got %d", found)
	}
}
