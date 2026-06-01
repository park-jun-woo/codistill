//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryResourceResponse 테스트
package laravel

import "testing"

func TestTryResourceResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function show() { return new UserResource($u); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryResourceResponse(t.TempDir(), ret, fi.src, nil, map[string]*fileInfo{})
	if resp == nil || resp.TypeName != "UserResource" {
		t.Fatalf("got %+v", resp)
	}
}
