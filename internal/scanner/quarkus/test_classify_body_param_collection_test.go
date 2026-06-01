//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyBodyParam_Collection 테스트
package quarkus

import "testing"

func TestClassifyBodyParam_Collection(t *testing.T) {
	// List<X> body → inner DTO type kept for resolution + array flag set.
	ep := &endpointInfo{}
	classifyBodyParam("List<CustomFieldJson>", "body", ep)
	if ep.bodyType != "CustomFieldJson" {
		t.Fatalf("List body type, got %q", ep.bodyType)
	}
	if !ep.bodyIsArray {
		t.Fatal("List body should set bodyIsArray")
	}

	// Set<X> behaves the same.
	ep2 := &endpointInfo{}
	classifyBodyParam("Set<Foo>", "body", ep2)
	if ep2.bodyType != "Foo" || !ep2.bodyIsArray {
		t.Fatalf("Set body, got %q array=%v", ep2.bodyType, ep2.bodyIsArray)
	}

	// Map<K,V> → free-form object: bodyType stays empty.
	ep3 := &endpointInfo{}
	classifyBodyParam("Map<String,Object>", "body", ep3)
	if ep3.bodyType != "" {
		t.Fatalf("Map body should be free-form (empty type), got %q", ep3.bodyType)
	}

	// Non-generic DTO body unchanged (regression guard).
	ep4 := &endpointInfo{}
	classifyBodyParam("FooDto", "body", ep4)
	if ep4.bodyType != "FooDto" || ep4.bodyIsArray {
		t.Fatalf("plain DTO body, got %q array=%v", ep4.bodyType, ep4.bodyIsArray)
	}
}
