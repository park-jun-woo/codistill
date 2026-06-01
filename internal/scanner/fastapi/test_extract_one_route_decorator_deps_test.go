//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractOneRoute_DecoratorDependencies 데코레이터 dependencies 통합 테스트
package fastapi

import "testing"

func TestExtractOneRoute_DecoratorDependencies(t *testing.T) {
	src := []byte(`
from fastapi import APIRouter, Depends

router = APIRouter(prefix="/admin")

@router.get("/", dependencies=[Depends(get_current_active_superuser)])
async def admin_list():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routes := extractRoutes(root, src, prefixes, nil, nil, "admin.py", nil)

	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if len(routes[0].middleware) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(routes[0].middleware))
	}
	if routes[0].middleware[0] != "get_current_active_superuser" {
		t.Errorf("expected get_current_active_superuser, got %s", routes[0].middleware[0])
	}
}
