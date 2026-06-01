//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExtractCallPath 테스트
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractCallPath(t *testing.T) {
	src := `package m
import "path"
const A = "/admin"
var _ = path.Join(A, "/login")
var _ = path.Join("/x", "/y", "/z")
var _ = path.Clean("/foo")
var dyn string
var _ = path.Join(dyn, "/login")
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}

	var calls []*ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			calls = append(calls, c)
		}
		return true
	})
	if len(calls) != 4 {
		t.Fatalf("want 4 calls, got %d", len(calls))
	}

	if got, ok := extractCallPath(info, calls[0]); !ok || got != "/admin/login" {
		t.Fatalf("const join: %q %v", got, ok)
	}
	if got, ok := extractCallPath(info, calls[1]); !ok || got != "/x/y/z" {
		t.Fatalf("multi join: %q %v", got, ok)
	}
	if got, ok := extractCallPath(info, calls[2]); !ok || got != "/foo" {
		t.Fatalf("clean: %q %v", got, ok)
	}
	if _, ok := extractCallPath(info, calls[3]); ok {
		t.Fatal("dynamic arg should not resolve")
	}
}
