//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findRouterAssignments 테스트
package fastapi

import "testing"

func TestFindRouterAssignments(t *testing.T) {
	src := []byte("app = FastAPI()\nrouter = APIRouter(prefix='/api')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	routers := findRouterAssignments(root, src, nil)
	if len(routers) < 1 {
		t.Fatalf("expected at least 1 router, got %d", len(routers))
	}
}
