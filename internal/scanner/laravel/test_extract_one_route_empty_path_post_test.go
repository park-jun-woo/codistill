//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_EmptyPathPost 테스트
package laravel

import "testing"

func TestExtractOneRoute_EmptyPathPost(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::post('', [AccountController::class, 'store']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	r := extractOneRoute(call, fi, "v1/accounts", nil)
	if r == nil || r.method != "POST" || r.path != "/v1/accounts" {
		t.Fatalf("expected POST /v1/accounts, got %+v", r)
	}
}
