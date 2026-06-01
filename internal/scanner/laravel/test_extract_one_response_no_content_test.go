//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneResponse_NoContent 테스트
package laravel

import "testing"

func TestExtractOneResponse_NoContent(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->noContent(); } }`)
	rets := findAllByType(fi.root, "return_statement")
	resp := extractOneResponse(t.TempDir(), rets[0], fi.src, nil, map[string]*fileInfo{})
	if resp == nil || resp.Status != "204" {
		t.Fatalf("got %+v", resp)
	}
}
