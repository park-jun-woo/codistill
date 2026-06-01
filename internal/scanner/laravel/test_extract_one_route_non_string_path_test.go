//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_NonStringPath 테스트 (회귀 가드)
package laravel

import "testing"

func TestExtractOneRoute_NonStringPath(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get($path, [AccountController::class, 'index']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractOneRoute(call, fi, "v1/accounts", nil); r != nil {
		t.Fatalf("expected nil for non-string path, got %+v", r)
	}
}
