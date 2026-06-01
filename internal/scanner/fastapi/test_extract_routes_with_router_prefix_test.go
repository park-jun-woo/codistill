//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractRoutes_WithRouterPrefix 테스트
package fastapi

import "testing"

func TestExtractRoutes_WithRouterPrefix(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, APIRouter

app = FastAPI()
router = APIRouter(prefix="/users")

@router.get("/")
async def list_users():
    pass

@router.get("/{user_id}")
async def get_user(user_id: int):
    pass

app.include_router(router)
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routes := extractRoutes(root, src, prefixes, nil, nil, "main.py", nil)

	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	if routes[0].path != "/users" {
		t.Fatalf("expected /users, got %s", routes[0].path)
	}
	if routes[1].path != "/users/{user_id}" {
		t.Fatalf("expected /users/{user_id}, got %s", routes[1].path)
	}
}
