//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveController_Path 테스트
package laravel

import "testing"

func TestResolveController_Path(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Http/Controllers/Api/PostController.php", `<?php class PostController {}`)
	if resolveController(dir, "PostController", nil, map[string]*fileInfo{}) == nil {
		t.Fatal("expected via Api path")
	}
}
