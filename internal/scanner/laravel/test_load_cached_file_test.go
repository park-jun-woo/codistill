//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestLoadCachedFile_Memoize 테스트 (miss 시 파싱·캐시 적재, hit 시 동일 포인터 반환)
package laravel

import "testing"

func TestLoadCachedFile_Memoize(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "modules/x/Foo.php", `<?php class Foo {}`)
	cache := map[string]*fileInfo{}

	fi := loadCachedFile(dir, "modules/x/Foo.php", cache)
	if fi == nil {
		t.Fatal("expected file parsed on miss")
	}
	if _, ok := cache["modules/x/Foo.php"]; !ok {
		t.Fatal("expected parsed file inserted into cache")
	}
	again := loadCachedFile(dir, "modules/x/Foo.php", cache)
	if again != fi {
		t.Fatal("expected cache hit to return the same fileInfo pointer")
	}
	if miss := loadCachedFile(dir, "modules/x/Missing.php", cache); miss != nil {
		t.Fatal("expected nil for nonexistent file")
	}
}
