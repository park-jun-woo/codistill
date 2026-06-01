//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestProviderRouteFilesRequire: Provider require base_path 로드 파일(서브디렉터리) 수집 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestProviderRouteFilesRequire(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		require base_path('routes/web/hardware.php');
	}
}
`)
	writeFile(t, dir, "routes/web/hardware.php", `<?php
Route::get('/hardware', [HardwareController::class, 'index']);
`)
	provFi, err := parseFile(dir, filepath.Join(dir, "app/Providers/RouteServiceProvider.php"))
	if err != nil {
		t.Fatal(err)
	}
	hwFi, err := parseFile(dir, filepath.Join(dir, "routes/web/hardware.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		"app/Providers/RouteServiceProvider.php": provFi,
		"routes/web/hardware.php":                hwFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/hardware" {
		t.Errorf("expected /hardware, got %q", routes[0].path)
	}
}
