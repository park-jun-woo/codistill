//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindFormRequestFile_PSR4 테스트
package laravel

import "testing"

func TestFindFormRequestFile_PSR4(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Requests/StoreReq.php", `<?php class StoreReq {}`)
	if findFormRequestFile(dir, "StoreReq", nil, map[string]*fileInfo{}) == nil {
		t.Fatal("expected to find via PSR-4 path")
	}
}
