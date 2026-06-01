//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestResolvePathArg 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolvePathArg(t *testing.T) {
	dir := t.TempDir()
	content := `package x;
class R {
  static final String P = "/api";
  static final String A="/1.0", B="/kb";
  static final String C=A+B;
  static final String UUID_PATTERN="[0-9a-f-]+";
  static final String SEARCH="search";
  static final String ANYTHING_PATTERN=".*";
  @Path(P) void a(){}
  @Path(C) void b(){}
  @Path("/{id:"+UUID_PATTERN+"}") void c(){}
  @Path("/"+SEARCH+"/{k:"+ANYTHING_PATTERN+"}") void d(){}
  @Path("/lit") void e(){}
  @Path(UNKNOWN_CONST) void f(){}
}`
	writeFile(t, dir, "R.java", content)
	abs := filepath.Join(dir, "R.java")
	src := []byte(content)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)

	want := map[string]string{
		"a": "/api",
		"b": "/1.0/kb",
		"c": "/{id}",
		"d": "/search/{k}",
		"e": "/lit",
		"f": "", // unresolved constant -> empty fallback, no panic
	}
	methods := findAllByType(root, "method_declaration")
	seen := map[string]bool{}
	for _, m := range methods {
		nameNode := findChildByType(m, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, src)
		w, ok := want[name]
		if !ok {
			continue
		}
		seen[name] = true
		ann := findAnnotation(m, src, AnnPath)
		var got string
		if ann != nil {
			got = resolvePathArg(ann, src, imports, abs, dir)
		}
		if got != w {
			t.Errorf("method %s: got %q want %q", name, got, w)
		}
	}
	for n := range want {
		if !seen[n] {
			t.Errorf("method %s not found in AST", n)
		}
	}
}
