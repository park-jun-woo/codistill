//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_Context 테스트
package quarkus

import "testing"

func TestClassifyParam_Context(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@Context HttpServletRequest req) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.bodyType != "" {
		t.Fatalf("@Context param should not be body, got %q", ep.bodyType)
	}
	if ep.formType != "" {
		t.Fatalf("@Context param should not be form, got %q", ep.formType)
	}
}
