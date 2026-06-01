//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseRouteDecorator 테스트
package fastapi

import "testing"

func TestParseRouteDecorator(t *testing.T) {
	src := []byte("@app.post('/items', status_code=201)\ndef create(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	method, routerVar, da := parseRouteDecorator(decs[0], src)
	if method != "POST" {
		t.Fatalf("method: got %q", method)
	}
	if da.path != "/items" {
		t.Fatalf("path: got %q", da.path)
	}
	if routerVar != "app" {
		t.Fatalf("routerVar: got %q", routerVar)
	}
	if da.statusCode != 201 {
		t.Fatalf("status: got %d", da.statusCode)
	}
}
