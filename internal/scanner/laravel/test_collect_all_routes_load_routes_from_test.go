//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesLoadRoutesFrom: $this->loadRoutesFrom(__DIR__.'/../Routes/api.php') + 그 파일의 Route::group(['prefix'=>'shop'], fn) 안 Route::get('/products') → /shop/products
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesLoadRoutesFrom(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "packages/Acme/Catalog/src/Providers/CatalogServiceProvider.php", `<?php
class CatalogServiceProvider {
	public function boot() {
		$this->loadRoutesFrom(__DIR__ . '/../Routes/api.php');
	}
}
`)
	writeFile(t, dir, "packages/Acme/Catalog/src/Routes/api.php", `<?php
Route::group(['prefix' => 'shop'], function () {
	Route::get('/products', [ProductController::class, 'index']);
});
`)
	provRel := "packages/Acme/Catalog/src/Providers/CatalogServiceProvider.php"
	routeRel := "packages/Acme/Catalog/src/Routes/api.php"
	provFi, err := parseFile(dir, filepath.Join(dir, provRel))
	if err != nil {
		t.Fatal(err)
	}
	routeFi, err := parseFile(dir, filepath.Join(dir, routeRel))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{provRel: provFi, routeRel: routeFi}

	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/shop/products" {
		t.Errorf("expected /shop/products, got %q", routes[0].path)
	}
}
