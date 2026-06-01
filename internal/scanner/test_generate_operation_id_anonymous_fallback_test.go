//ff:func feature=scan type=test control=sequence
//ff:what TestGenerateOperationID_AnonymousFallback 테스트
package scanner

import "testing"

func TestGenerateOperationID_AnonymousFallback(t *testing.T) {
	ep := Endpoint{Handler: "(anonymous)", Method: "GET", Path: "/api"}
	got := generateOperationID(ep)
	if got == "(anonymous)" {
		t.Fatalf("expected path/method fallback, got %q", got)
	}
	if got == "" {
		t.Fatal("expected non-empty fallback")
	}
}
