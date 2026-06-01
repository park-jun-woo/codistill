//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneResponse_JSON 테스트
package laravel

import "testing"

func TestExtractOneResponse_JSON(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json(['ok' => true], 201); } }`)
	rets := findAllByType(fi.root, "return_statement")
	if len(rets) == 0 {
		t.Fatal("no return")
	}
	resp := extractOneResponse(t.TempDir(), rets[0], fi.src, nil, map[string]*fileInfo{})
	if resp == nil || resp.Status != "201" {
		t.Fatalf("got %+v", resp)
	}
}
