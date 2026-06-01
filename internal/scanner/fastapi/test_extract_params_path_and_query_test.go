//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractParams_PathAndQuery 테스트
package fastapi

import "testing"

func TestExtractParams_PathAndQuery(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, Query

app = FastAPI()

@app.get("/users/{user_id}")
async def get_user(user_id: int, skip: int = 0, limit: int = Query(default=100)):
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routes := extractRoutes(root, src, prefixes, nil, nil, "main.py", nil)

	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	ri := routes[0]
	if len(ri.params) != 1 {
		t.Fatalf("expected 1 path param, got %d", len(ri.params))
	}
	if ri.params[0].Name != "user_id" {
		t.Fatalf("expected user_id, got %s", ri.params[0].Name)
	}
	if ri.params[0].Type != "integer" {
		t.Fatalf("expected integer, got %s", ri.params[0].Type)
	}
	if len(ri.query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ri.query))
	}
	if ri.query[0].Name != "skip" {
		t.Fatalf("expected skip, got %s", ri.query[0].Name)
	}
	if ri.query[1].Name != "limit" {
		t.Fatalf("expected limit, got %s", ri.query[1].Name)
	}
}
