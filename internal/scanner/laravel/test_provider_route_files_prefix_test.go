//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestProviderRouteFilesPrefix: Provider ->prefix('p')->group(base_path) prefix 전파 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestProviderRouteFilesPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		Route::prefix('p')->group(base_path('routes/y.php'));
	}
}
`)
	writeFile(t, dir, "routes/y.php", `<?php
Route::get('/thing', [ThingController::class, 'index']);
`)
	provFi, err := parseFile(dir, filepath.Join(dir, "app/Providers/RouteServiceProvider.php"))
	if err != nil {
		t.Fatal(err)
	}
	yFi, err := parseFile(dir, filepath.Join(dir, "routes/y.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		"app/Providers/RouteServiceProvider.php": provFi,
		"routes/y.php":                           yFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/p/thing" {
		t.Errorf("expected /p/thing, got %q", routes[0].path)
	}
}
