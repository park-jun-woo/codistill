//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestExtractOneRoute_RouterDeps 테스트
package fastapi

import "testing"

func TestExtractOneRoute_RouterDeps(t *testing.T) {
	src := []byte("@router.post('/items')\ndef create():\n    pass\n")
	root, _ := parsePython(src)
	defs := findAllByType(root, "decorated_definition")
	routerDeps := map[string][]string{"router": {"verify_token"}}
	ri := extractOneRoute(defs[0], src, map[string]string{"router": ""}, routerDeps, nil, "main.py", nil)
	if ri == nil {
		t.Fatal("expected route")
	}
	found := false
	for _, m := range ri.middleware {
		if m == "verify_token" {
			found = true
		}
	}
	if !found {
		t.Fatalf("router-level dep not applied: %v", ri.middleware)
	}
}
