//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what include_in_schema=True 및 비리터럴 값 라우트가 노출 유지되는지 검증한다
package fastapi

import "testing"

func TestExtractOneRoute_IncludeInSchemaTruePreserved(t *testing.T) {
	cases := []string{
		"@app.get('/shown', include_in_schema=True)\ndef h(): pass\n",
		"@app.get('/dyn', include_in_schema=settings.docs)\ndef h(): pass\n",
		"@app.get('/plain')\ndef h(): pass\n",
	}
	for _, src := range cases {
		root, err := parsePython([]byte(src))
		if err != nil {
			t.Fatal(err)
		}
		defs := findAllByType(root, "decorated_definition")
		if len(defs) == 0 {
			t.Fatalf("no decorated definition for %q", src)
		}
		ri := extractOneRoute(defs[0], []byte(src), nil, nil, nil, "main.py", nil)
		if ri == nil {
			t.Fatalf("expected route preserved for %q", src)
		}
	}
}
