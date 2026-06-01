//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestProviderRouteFilesGroup: Provider ->group(base_path) 로드 파일 수집 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestProviderRouteFilesGroup(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		Route::group(base_path('routes/admin.php'));
	}
}
`)
	writeFile(t, dir, "routes/admin.php", `<?php
Route::get('/x', [AdminController::class, 'index']);
`)
	provFi, err := parseFile(dir, filepath.Join(dir, "app/Providers/RouteServiceProvider.php"))
	if err != nil {
		t.Fatal(err)
	}
	adminFi, err := parseFile(dir, filepath.Join(dir, "routes/admin.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		"app/Providers/RouteServiceProvider.php": provFi,
		"routes/admin.php":                       adminFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/x" {
		t.Errorf("expected /x, got %q", routes[0].path)
	}
}
