//ff:func feature=scan type=test control=sequence topic=django
//ff:what include("a.b.urls")가 a/b/urls.py와 a/b/urls/__init__.py 양쪽 모듈 키에 매칭됨을 검증한다
package django

import "testing"

func TestResolveIncludeModulePackage(t *testing.T) {
	// relPathToModule maps both forms to the same dotted key "a.b.urls".
	if got := relPathToModule("a/b/urls.py"); got != "a.b.urls" {
		t.Fatalf("relPathToModule(file): got %q want a.b.urls", got)
	}
	if got := relPathToModule("a/b/urls/__init__.py"); got != "a.b.urls" {
		t.Fatalf("relPathToModule(pkg): got %q want a.b.urls", got)
	}
	// Module-file form.
	fileMap := map[string][]urlEntry{"a.b.urls": nil}
	if got, ok := resolveIncludeModule("a.b.urls", fileMap); !ok || got != "a.b.urls" {
		t.Errorf("file form: got %q ok=%v", got, ok)
	}
	// Package-init form resolves to the identical key.
	pkgMap := map[string][]urlEntry{"a.b.urls": nil}
	if got, ok := resolveIncludeModule("a.b.urls", pkgMap); !ok || got != "a.b.urls" {
		t.Errorf("package form: got %q ok=%v", got, ok)
	}
}
