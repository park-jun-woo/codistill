//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryCollectionResponse_None 테스트
package laravel

import "testing"

func TestTryCollectionResponse_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() { return $x; } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if resp := tryCollectionResponse(t.TempDir(), ret, fi.src, nil, map[string]*fileInfo{}); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}
