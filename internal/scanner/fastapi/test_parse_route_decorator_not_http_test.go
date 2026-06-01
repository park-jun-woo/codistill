//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseRouteDecorator_NotHTTP 테스트
package fastapi

import "testing"

func TestParseRouteDecorator_NotHTTP(t *testing.T) {
	src := []byte("@app.middleware('http')\ndef h(): pass\n")
	root, _ := parsePython(src)
	decs := findAllByType(root, "decorator")
	method, _, _ := parseRouteDecorator(decs[0], src)
	if method != "" {
		t.Fatalf("expected empty method, got %q", method)
	}
}
