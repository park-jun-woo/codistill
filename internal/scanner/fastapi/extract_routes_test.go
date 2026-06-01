//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractRoutes_Basic 테스트
package fastapi

import "testing"

func TestExtractRoutes_Basic(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "hello"}

@app.post("/items", status_code=201)
async def create_item():
    pass
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
	if routes[0].method != "GET" {
		t.Fatalf("expected GET, got %s", routes[0].method)
	}
	if routes[0].path != "/" {
		t.Fatalf("expected /, got %s", routes[0].path)
	}
	if routes[0].handler != "root" {
		t.Fatalf("expected root, got %s", routes[0].handler)
	}
	if routes[1].method != "POST" {
		t.Fatalf("expected POST, got %s", routes[1].method)
	}
	if routes[1].path != "/items" {
		t.Fatalf("expected /items, got %s", routes[1].path)
	}
	if routes[1].statusCode != 201 {
		t.Fatalf("expected status_code=201, got %d", routes[1].statusCode)
	}
}
