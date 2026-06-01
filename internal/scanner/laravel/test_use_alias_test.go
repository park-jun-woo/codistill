//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestUseAlias 테스트
package laravel

import "testing"

func TestUseAlias(t *testing.T) {
	fi := mustParsePHP(t, `<?php
use App\Foo as Bar;
use App\Baz;
`)
	clauses := findAllByType(fi.root, "namespace_use_clause")
	if len(clauses) != 2 {
		t.Fatalf("expected 2 use clauses, got %d", len(clauses))
	}
	if a := useAlias(clauses[0], fi.src); a != "Bar" {
		t.Errorf("aliased clause = %q, want Bar", a)
	}
	if a := useAlias(clauses[1], fi.src); a != "" {
		t.Errorf("non-aliased clause = %q, want empty", a)
	}
}
