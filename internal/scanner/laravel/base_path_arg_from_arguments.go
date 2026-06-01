//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what arguments 노드의 인자 중 base_path('X') 호출을 찾아 문자열 인자를 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// basePathArgFromArguments scans the arguments of a call for a base_path('X')
// argument and returns its literal string. Used for ->group(base_path('X')).
func basePathArgFromArguments(args *sitter.Node, src []byte) (string, bool) {
	for _, arg := range childrenOfType(args, "argument") {
		call := findChildByType(arg, "function_call_expression")
		if call == nil {
			continue
		}
		if rel, ok := extractBasePathArg(call, src); ok {
			return rel, true
		}
	}
	return "", false
}
