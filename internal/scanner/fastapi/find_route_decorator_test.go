//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findRouteDecorator 테스트
package fastapi

import "testing"

func TestFindRouteDecorator(t *testing.T) {
	src := []byte("@router.get('/users')\nasync def list_users(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := childrenOfType(root.Child(0), "decorator")
	if len(decs) == 0 {
		decs = findAllByType(root, "decorator")
	}
	if len(decs) == 0 {
		t.Fatal("no decorators")
	}
	method, routerVar, da := findRouteDecorator(decs, src)
	if method != "GET" || da.path != "/users" || routerVar != "router" {
		t.Fatalf("got method=%q path=%q var=%q", method, da.path, routerVar)
	}

	// empty list
	m2, _, _ := findRouteDecorator(nil, src)
	if m2 != "" {
		t.Fatal("expected empty for nil decorators")
	}
}
