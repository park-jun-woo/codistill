//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestFindRouteSourceFiles 테스트 (라우트 소스만 수집, 무관 PHP 제외)
package laravel

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindRouteSourceFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php`)
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php`)
	writeFile(t, dir, "packages/X/src/Routes/api.php", `<?php`)
	writeFile(t, dir, "app/Http/Controllers/UserController.php", `<?php`)
	writeFile(t, dir, "app/Models/User.php", `<?php`)
	writeFile(t, dir, "vendor/foo/routes/web.php", `<?php`)

	files, err := findRouteSourceFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		got[strings.ReplaceAll(rel, "\\", "/")] = true
	}
	want := []string{"routes/api.php", "app/Providers/RouteServiceProvider.php", "packages/X/src/Routes/api.php"}
	for _, w := range want {
		if !got[w] {
			t.Errorf("expected %q to be collected", w)
		}
	}
	if got["app/Http/Controllers/UserController.php"] {
		t.Error("controller must not be collected in stage 1")
	}
	if got["app/Models/User.php"] {
		t.Error("model must not be collected in stage 1")
	}
	if got["vendor/foo/routes/web.php"] {
		t.Error("vendor must be skipped")
	}
}
