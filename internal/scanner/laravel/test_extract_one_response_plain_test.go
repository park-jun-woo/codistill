//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneResponse_Plain 테스트
package laravel

import "testing"

func TestExtractOneResponse_Plain(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return $x; } }`)
	rets := findAllByType(fi.root, "return_statement")
	if resp := extractOneResponse(t.TempDir(), rets[0], fi.src, nil, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil for plain return, got %+v", resp)
	}
}
