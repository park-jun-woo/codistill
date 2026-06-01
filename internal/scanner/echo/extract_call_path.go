//ff:func feature=scan type=extract control=selection
//ff:what path.Join/filepath.Join/path.Clean 합성 호출에서 경로를 해석한다
package echo

import (
	"go/ast"
	"go/types"
)

// extractCallPath resolves a path from a whitelisted path-composition call
// (path.Join / filepath.Join / path.Clean). Each argument is resolved via the
// existing extractPathString machinery (literal / "+" concat / const ident).
// If any argument is unresolvable, it returns ("", false) so the route is dropped.
func extractCallPath(info *types.Info, call *ast.CallExpr) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	recv := identName(sel.X)
	if recv != "path" && recv != "filepath" {
		return "", false
	}
	switch sel.Sel.Name {
	case "Join":
		return joinCallArgs(info, call.Args)
	case "Clean":
		if len(call.Args) != 1 {
			return "", false
		}
		return extractPathString(info, call.Args[0])
	}
	return "", false
}
