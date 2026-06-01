//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveClassViaPSR4_NilSource 테스트
package laravel

import "testing"

func TestResolveClassViaPSR4_NilSource(t *testing.T) {
	if resolveClassViaPSR4(t.TempDir(), "X", nil, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil when source file is nil")
	}
}
