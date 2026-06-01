//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutesProviderWebhooksPrefix: Provider ->prefix('webhooks')->group(webhooks.php) prefix가 부착된다
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutesProviderWebhooksPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Providers/RouteServiceProvider.php", `<?php
class RouteServiceProvider {
	public function map() {
		Route::prefix('webhooks')->group(base_path('routes/webhooks.php'));
	}
}
`)
	writeFile(t, dir, "routes/webhooks.php", `<?php
Route::get('/source/github/redirect', [GithubController::class, 'redirect']);
`)
	provFi, err := parseFile(dir, filepath.Join(dir, "app/Providers/RouteServiceProvider.php"))
	if err != nil {
		t.Fatal(err)
	}
	whFi, err := parseFile(dir, filepath.Join(dir, "routes/webhooks.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{
		"app/Providers/RouteServiceProvider.php": provFi,
		"routes/webhooks.php":                    whFi,
	}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/webhooks/source/github/redirect" {
		t.Errorf("expected /webhooks/source/github/redirect, got %q", routes[0].path)
	}
}
