//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what include_router(child, include_in_schema=False)가 child 라우터를 hidden으로 표시하는지 검증한다
package fastapi

import "testing"

func TestResolveHiddenRouters_IncludeRouterHidden(t *testing.T) {
	src := []byte("app = FastAPI()\nchild = APIRouter()\napp.include_router(child, include_in_schema=False)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	hidden := resolveHiddenRouters(root, src, nil)
	if !hidden["child"] {
		t.Fatalf("expected child marked hidden, got %v", hidden)
	}
	if hidden["app"] {
		t.Fatal("app should not be hidden")
	}
}
