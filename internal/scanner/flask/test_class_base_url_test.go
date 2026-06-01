//ff:func feature=scan type=test control=sequence topic=flask
//ff:what classBaseURL이 명시 base_url과 기본 컨벤션을 해석하는지 검증한다
package flask

import "testing"

func TestClassBaseURL(t *testing.T) {
	// Explicit base_url wins.
	root, src := mustParse(t, "class BarApi(ModelRestApi):\n    base_url = \"/custom/bar\"\n")
	cls := findChildByType(root, "class_definition")
	if got := classBaseURL(cls, "BarApi", src); got != "/custom/bar" {
		t.Fatalf("explicit base_url: got %q", got)
	}

	// route_base drives the /api/v1 default.
	root, src = mustParse(t, "class BarApi(ModelRestApi):\n    route_base = \"bar2\"\n")
	cls = findChildByType(root, "class_definition")
	if got := classBaseURL(cls, "BarApi", src); got != "/api/v1/bar2" {
		t.Fatalf("route_base default: got %q", got)
	}

	// No attributes: class-name fallback.
	root, src = mustParse(t, "class FooApi(BaseApi):\n    pass\n")
	cls = findChildByType(root, "class_definition")
	if got := classBaseURL(cls, "FooApi", src); got != "/api/v1/foo" {
		t.Fatalf("class-name fallback: got %q", got)
	}
}
