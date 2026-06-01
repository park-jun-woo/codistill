//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestObjectPatternHasName 테스트
package express

import "testing"

func TestObjectPatternHasName(t *testing.T) {
	fi := mustParse(t, []byte(`const {BASE_API_PATH} = require('../url');`))
	pats := findAllByType(fi.Root, "object_pattern")
	if len(pats) == 0 {
		t.Fatal("no object_pattern")
	}
	pat := pats[0]
	if !objectPatternHasName(pat, fi.Src, "BASE_API_PATH") {
		t.Fatal("expected name present")
	}
	if objectPatternHasName(pat, fi.Src, "OTHER") {
		t.Fatal("unexpected name match")
	}
}
