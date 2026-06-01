//ff:func feature=scan type=test control=sequence topic=flask
//ff:what classHTTPMethods가 HTTP 메서드만 추출하고 헬퍼는 제외하는지 검증한다
package flask

import "testing"

func TestClassHTTPMethods_ExcludesHelpers(t *testing.T) {
	src := []byte(`class UserAPI(Resource):
    def get(self):
        return {}
    def post(self):
        return {}
    def __helper(self):
        pass
    def utility(self):
        pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_definition")
	if len(classes) != 1 {
		t.Fatalf("expected 1 class, got %d", len(classes))
	}
	methods := classHTTPMethods(classes[0], src)
	if len(methods) != 2 {
		t.Fatalf("expected 2 HTTP methods, got %d", len(methods))
	}
	if methods[0].name != "GET" {
		t.Errorf("expected GET, got %s", methods[0].name)
	}
	if methods[1].name != "POST" {
		t.Errorf("expected POST, got %s", methods[1].name)
	}
}
