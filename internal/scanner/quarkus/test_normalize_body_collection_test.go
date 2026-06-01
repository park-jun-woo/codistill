//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestNormalizeBodyCollection 테스트
package quarkus

import "testing"

func TestNormalizeBodyCollection(t *testing.T) {
	if inner, kind := normalizeBodyCollection("List<CustomFieldJson>"); kind != bodyCollectionArray || inner != "CustomFieldJson" {
		t.Fatalf("List: got %q kind=%d", inner, kind)
	}
	if inner, kind := normalizeBodyCollection("Set<Foo>"); kind != bodyCollectionArray || inner != "Foo" {
		t.Fatalf("Set: got %q kind=%d", inner, kind)
	}
	if inner, kind := normalizeBodyCollection("Collection<Bar>"); kind != bodyCollectionArray || inner != "Bar" {
		t.Fatalf("Collection: got %q kind=%d", inner, kind)
	}
	if inner, kind := normalizeBodyCollection("Iterable<Baz>"); kind != bodyCollectionArray || inner != "Baz" {
		t.Fatalf("Iterable: got %q kind=%d", inner, kind)
	}
	if _, kind := normalizeBodyCollection("Map<String,Object>"); kind != bodyCollectionMap {
		t.Fatalf("Map should be map kind, got %d", kind)
	}
	if _, kind := normalizeBodyCollection("HashMap<String,Foo>"); kind != bodyCollectionMap {
		t.Fatalf("HashMap should be map kind, got %d", kind)
	}
	if inner, kind := normalizeBodyCollection("FooDto"); kind != bodyCollectionNone || inner != "FooDto" {
		t.Fatalf("plain DTO should be none, got %q kind=%d", inner, kind)
	}
}
