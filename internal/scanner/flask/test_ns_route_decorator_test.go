//ff:func feature=scan type=test control=sequence topic=flask
//ff:what nsRouteDecorator가 @ns.route 데코레이터에서 ns 변수와 path를 추출하고 비route는 거르는지 검증한다
package flask

import "testing"

func TestNSRouteDecorator(t *testing.T) {
	src := []byte(`@inner_api_ns.route("/invoke/llm")
@ns.expect(model)
class X(Resource):
    def post(self):
        pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	dd := findAllByType(root, "decorated_definition")[0]
	decs := childrenOfType(dd, "decorator")

	nsVar, path, ok := nsRouteDecorator(decs[0], src)
	if !ok || nsVar != "inner_api_ns" || path != "/invoke/llm" {
		t.Errorf("route decorator: ok=%v nsVar=%q path=%q", ok, nsVar, path)
	}
	if _, _, ok := nsRouteDecorator(decs[1], src); ok {
		t.Errorf("expect decorator should not be a route")
	}
}
