//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what include_in_schema=False 라우터의 모든 라우트가 스킵되는지 검증한다
package fastapi

import "testing"

func TestExtractRoutes_HiddenRouter(t *testing.T) {
	src := []byte("r = APIRouter(include_in_schema=False)\n@r.get('/y')\ndef y(): pass\n@r.post('/z')\ndef z(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	hidden := resolveHiddenRouters(root, src, nil)
	if !hidden["r"] {
		t.Fatalf("expected router r marked hidden, got %v", hidden)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	routes := extractRoutes(root, src, prefixes, nil, hidden, "main.py", nil)
	if len(routes) != 0 {
		t.Fatalf("expected 0 routes from hidden router, got %d", len(routes))
	}
}
