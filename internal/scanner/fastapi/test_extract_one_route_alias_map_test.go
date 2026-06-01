//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractOneRoute_AliasMapMiddleware 별칭 맵 미들웨어 통합 테스트
package fastapi

import "testing"

func TestExtractOneRoute_AliasMapMiddleware(t *testing.T) {
	src := []byte(`
from fastapi import APIRouter

router = APIRouter()

@router.get("/me")
async def read_user_me(current_user: CurrentUser):
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	aliasMap := map[string]string{
		"CurrentUser": "get_current_user",
		"SessionDep":  "get_db",
	}
	routes := extractRoutes(root, src, prefixes, nil, nil, "users.py", aliasMap)

	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if len(routes[0].middleware) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(routes[0].middleware))
	}
	if routes[0].middleware[0] != "get_current_user" {
		t.Errorf("expected get_current_user, got %s", routes[0].middleware[0])
	}
	if routes[0].bodyType != "" {
		t.Errorf("body type should be empty, got %s", routes[0].bodyType)
	}
}
