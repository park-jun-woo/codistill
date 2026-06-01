//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractRoutes_RouterDeps 라우터 생성자 dependencies 전파 통합 테스트
package fastapi

import "testing"

func TestExtractRoutes_RouterDeps(t *testing.T) {
	src := []byte(`
from fastapi import APIRouter, Depends

router = APIRouter(prefix="/sneakers", dependencies=[Depends(verify_token)])

@router.get("/")
async def list_sneakers():
    pass

@router.post("/", dependencies=[Depends(is_admin)])
async def create_sneaker():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routerDeps := resolveRouterDeps(root, src, nil)
	routes := extractRoutes(root, src, prefixes, routerDeps, nil, "sneakers.py", nil)

	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}

	// First endpoint: only router-level dep
	r0 := routes[0]
	if len(r0.middleware) != 1 {
		t.Fatalf("route 0: expected 1 middleware, got %d: %v", len(r0.middleware), r0.middleware)
	}
	if r0.middleware[0] != "verify_token" {
		t.Errorf("route 0: expected verify_token, got %s", r0.middleware[0])
	}

	// Second endpoint: router-level dep + decorator dep
	r1 := routes[1]
	if len(r1.middleware) != 2 {
		t.Fatalf("route 1: expected 2 middleware, got %d: %v", len(r1.middleware), r1.middleware)
	}
	if r1.middleware[0] != "verify_token" {
		t.Errorf("route 1: expected verify_token first, got %s", r1.middleware[0])
	}
	if r1.middleware[1] != "is_admin" {
		t.Errorf("route 1: expected is_admin second, got %s", r1.middleware[1])
	}
}
