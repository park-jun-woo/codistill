//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindResourceFile_NotFound 테스트
package laravel

import "testing"

func TestFindResourceFile_NotFound(t *testing.T) {
	if findResourceFile(t.TempDir(), "Missing", nil, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}
