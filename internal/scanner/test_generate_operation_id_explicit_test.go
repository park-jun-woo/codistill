//ff:func feature=scan type=test control=sequence
//ff:what TestGenerateOperationID_Explicit 테스트
package scanner

import "testing"

func TestGenerateOperationID_Explicit(t *testing.T) {
	ep := Endpoint{OperationID: "healthCheck", Handler: "(anonymous)", Method: "GET", Path: "/api"}
	got := generateOperationID(ep)
	if got != "healthCheck" {
		t.Fatalf("expected explicit operationId healthCheck, got %q", got)
	}
}
