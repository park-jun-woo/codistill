//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveClassViaPSR4 테스트 (use 임포트 + composer psr-4로 비표준 위치 클래스 해석)
package laravel

import "testing"

func TestResolveClassViaPSR4(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "composer.json", `{"autoload":{"psr-4":{"Acme\\":"lib/"}}}`)
	writeFile(t, dir, "lib/Controllers/WidgetController.php", `<?php
namespace Acme\Controllers;
class WidgetController {}
`)
	src := mustParsePHP(t, `<?php
use Acme\Controllers\WidgetController;
`)
	cache := map[string]*fileInfo{}
	fi := resolveClassViaPSR4(dir, "WidgetController", &src, cache)
	if fi == nil {
		t.Fatal("expected to resolve WidgetController via PSR-4")
	}
	if !classMatches(fi, "WidgetController") {
		t.Fatal("resolved file does not contain WidgetController")
	}
}
