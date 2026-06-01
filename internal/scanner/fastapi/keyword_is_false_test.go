//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what keywordIsFalse 테스트
package fastapi

import "testing"

func TestKeywordIsFalse(t *testing.T) {
	cases := []struct {
		src  string
		want bool
	}{
		{"@app.get('/a', include_in_schema=False)\ndef f(): pass\n", true},
		{"@app.get('/a', include_in_schema=True)\ndef f(): pass\n", false},
		{"@app.get('/a', include_in_schema=settings.x)\ndef f(): pass\n", false},
		{"@app.get('/a')\ndef f(): pass\n", false},
	}
	for _, c := range cases {
		root, err := parsePython([]byte(c.src))
		if err != nil {
			t.Fatal(err)
		}
		decs := findAllByType(root, "decorator")
		if len(decs) == 0 {
			t.Fatalf("no decorator for %q", c.src)
		}
		call := findChildByType(decs[0], "call")
		args := findChildByType(call, "argument_list")
		got := keywordIsFalse(args, "include_in_schema", []byte(c.src))
		if got != c.want {
			t.Fatalf("src=%q got=%v want=%v", c.src, got, c.want)
		}
	}
}
