//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryCollectionResponse 테스트
package laravel

import "testing"

func TestTryCollectionResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function index() {
		return UserResource::collection($users);
	} }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryCollectionResponse(t.TempDir(), ret, fi.src, nil, map[string]*fileInfo{})
	if resp == nil || resp.TypeName != "[]UserResource" {
		t.Fatalf("got %+v", resp)
	}
}
