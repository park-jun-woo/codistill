//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesNoPackage: 패키지 Provider가 없는 앱은 동작 불변(루트 라우트만 폴백 수집, 추가 0)
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesNoPackage(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
`)
	apiFi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{"routes/api.php": apiFi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/api/users" {
		t.Errorf("expected /api/users, got %q", routes[0].path)
	}
}
