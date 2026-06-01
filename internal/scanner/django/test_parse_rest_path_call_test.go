//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseRestPathCall_MethodKeywords 테스트
package django

import "testing"

func TestParseRestPathCall_MethodKeywords(t *testing.T) {
	src := []byte("urlpatterns = [rest_path('realm', GET=get_realm, PATCH=update_realm)]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := djFirst(t, root, "call")
	entry := parsePathCall(call, src)
	if entry == nil {
		t.Fatal("expected url entry")
	}
	if entry.pattern != "realm" {
		t.Errorf("pattern: %q", entry.pattern)
	}
	if got := entry.methodViews["GET"]; got != "get_realm" {
		t.Errorf("GET view: %q", got)
	}
	if got := entry.methodViews["PATCH"]; got != "update_realm" {
		t.Errorf("PATCH view: %q", got)
	}
	if len(entry.methodViews) != 2 {
		t.Errorf("methodViews count: %d", len(entry.methodViews))
	}
}
