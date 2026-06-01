//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindFormRequestFile_NotFound 테스트
package laravel

import "testing"

func TestFindFormRequestFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	if findFormRequestFile(dir, "Missing", nil, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}
