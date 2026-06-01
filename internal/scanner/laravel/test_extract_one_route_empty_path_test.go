//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_EmptyPath 테스트
package laravel

import "testing"

func TestExtractOneRoute_EmptyPath(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('', [AccountController::class, 'index']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	r := extractOneRoute(call, fi, "v1/accounts", nil)
	if r == nil || r.method != "GET" || r.path != "/v1/accounts" {
		t.Fatalf("expected GET /v1/accounts, got %+v", r)
	}
}
