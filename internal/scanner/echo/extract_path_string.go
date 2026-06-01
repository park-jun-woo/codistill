//ff:func feature=scan type=extract control=selection
//ff:what 리터럴·문자열 연결·상수 식별자에서 경로를 추출한다
package echo

import (
	"go/ast"
	"go/token"
	"go/types"
)

// extractPathString extracts a path from a literal, string concatenation, or const identifier.
func extractPathString(info *types.Info, expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return unquote(e.Value), true
		}
	case *ast.BinaryExpr:
		return extractBinaryPath(info, e)
	case *ast.Ident, *ast.SelectorExpr:
		if v := resolveExprConst(info, e); v != "" {
			return v, true
		}
	case *ast.CallExpr:
		return extractCallPath(info, e)
	}
	return "", false
}
