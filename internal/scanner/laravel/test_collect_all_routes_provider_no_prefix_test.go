//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesProviderNoPrefix: Provider가 prefix 없이 api.php 로드 시 /api/api 이중부착이 없다
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesProviderNoPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		Route::middleware('api')->group(base_path('routes/api.php'));
	}
}
`)
	writeFile(t, dir, "routes/api.php", `<?php
Route::group(['prefix' => 'api'], function () {
	Route::get('/f/inbox', [InboxController::class, 'index']);
});
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
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/api/f/inbox" {
		t.Errorf("expected /api/f/inbox (single prefix), got %q", routes[0].path)
	}
}
