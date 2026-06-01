//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseRouteArg 테스트
package actix

import "testing"

func TestParseRouteArg(t *testing.T) {
	src := []byte(`fn f() { x.route(web::get().to(get_user)); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if call == nil {
		t.Fatal("no .route call")
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	method, handler, pathOverride := parseRouteArg(args, src)
	if method != "GET" {
		t.Errorf("method = %q, want GET", method)
	}
	if handler != "get_user" {
		t.Errorf("handler = %q, want get_user", handler)
	}
	if pathOverride != "" {
		t.Errorf("pathOverride = %q, want empty (A-form)", pathOverride)
	}
}
