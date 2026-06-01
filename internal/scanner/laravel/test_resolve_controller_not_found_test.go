//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveController_NotFound 테스트
package laravel

import "testing"

func TestResolveController_NotFound(t *testing.T) {
	if resolveController(t.TempDir(), "Missing", nil, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil")
	}
}
