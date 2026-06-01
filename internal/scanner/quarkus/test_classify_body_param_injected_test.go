//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyBodyParam_Injected 테스트
package quarkus

import "testing"

func TestClassifyBodyParam_Injected(t *testing.T) {
	ep := &endpointInfo{}
	classifyBodyParam("HttpServletRequest", "req", ep)
	if ep.bodyType != "" {
		t.Fatalf("injected type should be excluded from body, got %q", ep.bodyType)
	}

	ep2 := &endpointInfo{}
	classifyBodyParam("UriInfo", "uriInfo", ep2)
	if ep2.bodyType != "" {
		t.Fatalf("injected type should be excluded from body, got %q", ep2.bodyType)
	}

	// 정상 POJO 본문은 그대로 body여야 한다(회귀 방지).
	ep3 := &endpointInfo{}
	classifyBodyParam("FooDto", "foo", ep3)
	if ep3.bodyType != "FooDto" {
		t.Fatalf("normal DTO should remain body, got %q", ep3.bodyType)
	}
}
