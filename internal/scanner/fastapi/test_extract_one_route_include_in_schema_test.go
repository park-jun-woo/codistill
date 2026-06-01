//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what include_in_schema=False 데코레이터 라우트가 스킵되는지 검증한다
package fastapi

import "testing"

func TestExtractOneRoute_IncludeInSchemaFalse(t *testing.T) {
	src := []byte("@app.get('/hidden', include_in_schema=False)\ndef h(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	if len(defs) == 0 {
		t.Fatal("no decorated definition")
	}
	ri := extractOneRoute(defs[0], src, nil, nil, nil, "main.py", nil)
	if ri != nil {
		t.Fatalf("expected nil for include_in_schema=False route, got %+v", ri)
	}
}
