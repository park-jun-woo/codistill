//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what path.Join 가변 인자를 개별 해석 후 JoinPath 좌측 fold로 결합한다
package echo

import (
	"go/ast"
	"go/types"
)

// joinCallArgs resolves each argument via extractPathString and folds them with
// scanner.JoinPath (left-to-right). Any unresolvable arg yields ("", false).
func joinCallArgs(info *types.Info, args []ast.Expr) (string, bool) {
	if len(args) == 0 {
		return "", false
	}
	var joined string
	for i, arg := range args {
		part, ok := extractPathString(info, arg)
		if !ok {
			return "", false
		}
		joined = foldPart(joined, part, i)
	}
	return joined, true
}
