//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScanMacroEndpoints 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanMacroEndpoints(t *testing.T) {
	root, err := parseRust([]byte(macroRoutesSource))
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: []byte(macroRoutesSource), root: root}
	files := []*fileInfo{fi}

	sIdx := buildStructIndex(files)
	cache := map[string][]scanner.Field{}
	handlerFuncs := map[string]*handlerInfo{}

	registeredHandlers := collectRegisteredHandlers(files)
	eps := scanMacroEndpoints(files, sIdx, cache, handlerFuncs, registeredHandlers)
	if len(eps) != 4 {
		t.Fatalf("expected 4 macro endpoints, got %d: %+v", len(eps), eps)
	}

	if len(handlerFuncs) == 0 {
		t.Error("expected handlerFuncs to be populated")
	}
}
