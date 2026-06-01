//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScanMacroEndpoints_NoFiles 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanMacroEndpoints_NoFiles(t *testing.T) {
	eps := scanMacroEndpoints(nil, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{}, map[string]bool{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %+v", eps)
	}
}
