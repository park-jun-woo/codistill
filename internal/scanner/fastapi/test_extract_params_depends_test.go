//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractParams_Depends 테스트
package fastapi

import "testing"

func TestExtractParams_Depends(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, Depends

app = FastAPI()

@app.get("/protected")
async def protected(current_user: str = Depends(get_current_user)):
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
	if len(routes[0].middleware) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(routes[0].middleware))
	}
	if routes[0].middleware[0] != "get_current_user" {
		t.Fatalf("expected get_current_user, got %s", routes[0].middleware[0])
	}
}
