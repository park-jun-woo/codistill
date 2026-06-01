//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesProviderDedup: Provider가 명시 로드한 api.php는 폴백과 중복 수집되지 않는다
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesProviderDedup(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		Route::prefix('api')->group(base_path('routes/api.php'));
	}
}
`)
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
`)
	provFi, err := parseFile(dir, filepath.Join(dir, "app/Providers/RouteServiceProvider.php"))
	if err != nil {
		t.Fatal(err)
	}
	apiFi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		"app/Providers/RouteServiceProvider.php": provFi,
		"routes/api.php":                         apiFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route (no duplicate), got %d", len(routes))
	}
	if routes[0].path != "/api/users" {
		t.Errorf("expected /api/users, got %q", routes[0].path)
	}
}
