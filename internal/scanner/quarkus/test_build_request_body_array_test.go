//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildRequest_BodyArray 테스트
package quarkus

import "testing"

func TestBuildRequest_BodyArray(t *testing.T) {
	// An array body re-applies the slice marker so the shared schema builder
	// emits type:array + items:$ref(X) (no "<"/">" leaks into schema keys).
	r := buildRequest(endpointInfo{bodyType: "CustomFieldJson", bodyVarName: "body", bodyIsArray: true})
	if r == nil || r.Body == nil {
		t.Fatalf("got %+v", r)
	}
	if r.Body.TypeName != "CustomFieldJson[]" {
		t.Fatalf("array body type name, got %q", r.Body.TypeName)
	}
}
