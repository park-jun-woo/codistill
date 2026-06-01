//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestCollectAllRoutesNestedRequire: Provider 로드 web.php의 Route::group(['prefix'=>'admin'], fn){ require 'sales-routes.php'; } 안 형제 파일 라우트가 /admin prefix로 수집된다
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesNestedRequire(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "packages/Acme/Admin/src/Providers/AdminServiceProvider.php", `<?php
class AdminServiceProvider {
	public function boot() {
		Route::middleware(['web'])->group(__DIR__ . '/../Routes/web.php');
	}
}
`)
	writeFile(t, dir, "packages/Acme/Admin/src/Routes/web.php", `<?php
require 'auth-routes.php';

Route::group(['prefix' => 'admin', 'middleware' => ['admin']], function () {
	require 'sales-routes.php';
});
`)
	writeFile(t, dir, "packages/Acme/Admin/src/Routes/auth-routes.php", `<?php
Route::get('/login', [AuthController::class, 'show']);
`)
	writeFile(t, dir, "packages/Acme/Admin/src/Routes/sales-routes.php", `<?php
Route::get('/orders', [OrderController::class, 'index']);
`)
	rels := []string{
		"packages/Acme/Admin/src/Providers/AdminServiceProvider.php",
		"packages/Acme/Admin/src/Routes/web.php",
		"packages/Acme/Admin/src/Routes/auth-routes.php",
		"packages/Acme/Admin/src/Routes/sales-routes.php",
	}
	parsed := map[string]*fileInfo{}
	for _, rel := range rels {
		fi, err := parseFile(dir, filepath.Join(dir, rel))
		if err != nil {
			t.Fatal(err)
		}
		parsed[rel] = fi
	}

	routes := collectAllRoutes(parsed)
	got := map[string]bool{}
	for _, r := range routes {
		got[r.path] = true
	}
	if !got["/login"] {
		t.Errorf("expected top-level require route /login; got %v", routes)
	}
	if !got["/admin/orders"] {
		t.Errorf("expected nested-group require route /admin/orders; got %v", routes)
	}
}
