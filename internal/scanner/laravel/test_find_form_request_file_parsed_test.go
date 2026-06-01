//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindFormRequestFile_Parsed 테스트
package laravel

import "testing"

func TestFindFormRequestFile_Parsed(t *testing.T) {
	fi := mustParsePHP(t, `<?php class StoreReq {}`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	if findFormRequestFile("/root", "StoreReq", nil, parsed) == nil {
		t.Fatal("expected to find via parsed files")
	}
}
