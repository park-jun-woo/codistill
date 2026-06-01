//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestImportMap 테스트
package laravel

import "testing"

func TestImportMap(t *testing.T) {
	fi := mustParsePHP(t, `<?php
use App\Http\Controllers\Api\UserController;
use App\Http\Controllers\PostController as PC;
`)
	m := importMap(&fi)
	if m["UserController"] != "App\\Http\\Controllers\\Api\\UserController" {
		t.Errorf("UserController = %q", m["UserController"])
	}
	if m["PC"] != "App\\Http\\Controllers\\PostController" {
		t.Errorf("alias PC = %q", m["PC"])
	}
	if _, ok := m["PostController"]; ok {
		t.Error("aliased class should be keyed by alias, not original short name")
	}
}
