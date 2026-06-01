//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryResourceResponse_None 테스트
package laravel

import "testing"

func TestTryResourceResponse_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show() { return $x; } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if resp := tryResourceResponse(t.TempDir(), ret, fi.src, nil, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}
