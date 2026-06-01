//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what classSuperclasses가 식별자/attribute 상위클래스를 정확히 식별하는지 검증한다
package flask

import (
	"strings"
	"testing"
)

func TestClassSuperclasses_Identifies(t *testing.T) {
	cases := map[string]string{
		"class A(Resource):\n    pass\n":                        "Resource",
		"class B(flask_restful.Resource):\n    pass\n":          "flask_restful.Resource",
		"class C(ModelRestApi):\n    pass\n":                    "ModelRestApi",
		"class D(flask_restful.Resource, BaseApi):\n    pass\n": "flask_restful.Resource,BaseApi",
	}
	for src, want := range cases {
		root, err := parsePython([]byte(src))
		if err != nil {
			t.Fatal(err)
		}
		cls := findAllByType(root, "class_definition")[0]
		got := strings.Join(classSuperclasses(cls, []byte(src)), ",")
		if got != want {
			t.Errorf("src %q: expected %q, got %q", src, want, got)
		}
	}
}
