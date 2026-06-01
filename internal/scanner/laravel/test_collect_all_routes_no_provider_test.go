//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesNoProvider: Provider 등록 없는 api.php 전용 앱은 기존 동작 유지(중복 0)
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesNoProvider(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{"routes/api.php": fi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route (no duplicate), got %d", len(routes))
	}
	if routes[0].path != "/api/users" {
		t.Errorf("expected /api/users, got %q", routes[0].path)
	}
}
