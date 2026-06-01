//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestClassifyParam_ContextMixed 테스트
package quarkus

import "testing"

func TestClassifyParam_ContextMixed(t *testing.T) {
	root, src := parseQ(t, `class R { void m(@Context UriInfo uriInfo, FooDto dto) {} }`)
	params := findAllByType(root, "formal_parameter")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	ep := &endpointInfo{}
	for _, p := range params {
		classifyParam(p, src, ep, nil, "", "")
	}
	if ep.bodyType != "FooDto" {
		t.Fatalf("only FooDto should be body, got %q", ep.bodyType)
	}
}
