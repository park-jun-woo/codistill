//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestDirRelativeRouteFile: __DIR__ . '/../Routes/web.php' 가 Provider 디렉터리 기준으로 정규화되어 parsedFiles 키와 매칭된다
package laravel

import (
	"path/filepath"
	"testing"
)

func TestDirRelativeRouteFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "packages/Webkul/Admin/src/Providers/AdminServiceProvider.php", `<?php
class AdminServiceProvider {
	public function boot() {
		Route::group(__DIR__ . '/../Routes/web.php');
	}
}
`)
	provRel := "packages/Webkul/Admin/src/Providers/AdminServiceProvider.php"
	provFi, err := parseFile(dir, filepath.Join(dir, provRel))
	if err != nil {
		t.Fatal(err)
	}
	refs := providerRouteFiles(map[string]*fileInfo{provRel: provFi})
	if len(refs) != 1 {
		t.Fatalf("expected 1 ref, got %d", len(refs))
	}
	want := "packages/Webkul/Admin/src/Routes/web.php"
	if refs[0].relPath != want {
		t.Errorf("expected relPath %q, got %q", want, refs[0].relPath)
	}
}
