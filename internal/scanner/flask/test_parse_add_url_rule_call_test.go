//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseAddURLRuleCall이 비대상 호출·path누락을 거르고 !선두1문자만 스트립하는지 검증한다
package flask

import "testing"

func TestParseAddURLRuleCall_RejectsAndStrips(t *testing.T) {
	// Non-add_url_rule call is rejected.
	call, src := firstCall(t, `api.add_resource(Foo, '/x')`)
	if _, ok := parseAddURLRuleCall(call, src); ok {
		t.Errorf("expected non-add_url_rule call to be rejected")
	}

	// add_url_rule with no string argument at all is rejected.
	call, src = firstCall(t, `bp.add_url_rule(rule_var, ep_var, RHX)`)
	if _, ok := parseAddURLRuleCall(call, src); ok {
		t.Errorf("expected add_url_rule without string path to be rejected")
	}

	// Only the leading "!" is stripped; an interior "!" is preserved.
	call, src = firstCall(t, `bp.add_url_rule('!/a!b', 'ep', RHX)`)
	reg, ok := parseAddURLRuleCall(call, src)
	if !ok {
		t.Fatalf("expected match")
	}
	if reg.rawPath != "/a!b" || !reg.appRoot {
		t.Errorf("leading ! strip wrong: %+v", reg)
	}
}
