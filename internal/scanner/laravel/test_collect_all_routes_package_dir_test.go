//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesPackageDir: 패키지 Provider의 ->group(__DIR__ . '/../Routes/api.php') 로드 + prefix 전파 수집
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesPackageDir(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "packages/Webkul/Shop/src/Providers/ShopServiceProvider.php", `<?php
class ShopServiceProvider {
	public function boot() {
		Route::prefix('shop')->middleware(['web'])->group(__DIR__ . '/../Routes/api.php');
	}
}
`)
	writeFile(t, dir, "packages/Webkul/Shop/src/Routes/api.php", `<?php
Route::get('/y', [ShopController::class, 'index']);
`)
	provRel := "packages/Webkul/Shop/src/Providers/ShopServiceProvider.php"
	routeRel := "packages/Webkul/Shop/src/Routes/api.php"
	provFi, err := parseFile(dir, filepath.Join(dir, provRel))
	if err != nil {
		t.Fatal(err)
	}
	routeFi, err := parseFile(dir, filepath.Join(dir, routeRel))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		provRel:  provFi,
		routeRel: routeFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/shop/y" {
		t.Errorf("expected /shop/y, got %q", routes[0].path)
	}
}
