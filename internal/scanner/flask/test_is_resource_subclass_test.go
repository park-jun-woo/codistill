//ff:func feature=scan type=test control=sequence topic=flask
//ff:what isResourceSubclass가 직접/dotted/alias 상속을 판정하는지 검증한다
package flask

import "testing"

func TestIsResourceSubclass_Variants(t *testing.T) {
	if !isResourceSubclass([]string{"Resource"}, importAlias{}) {
		t.Error("direct Resource should match")
	}
	if !isResourceSubclass([]string{"flask_restful.Resource"}, importAlias{}) {
		t.Error("dotted flask_restful.Resource should match")
	}
	if !isResourceSubclass([]string{"R"}, importAlias{"R": "Resource"}) {
		t.Error("aliased Resource should match")
	}
	if isResourceSubclass([]string{"BaseApi", "object"}, importAlias{}) {
		t.Error("non-Resource superclasses should not match")
	}
}
