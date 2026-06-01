//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what function_call_expression이 base_path('X')이면 그 문자열 인자를 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractBasePathArg returns the literal string argument of a base_path('X')
// function call. It returns ok=false for any other call or a non-static
// argument (variable/concatenation), which the caller skips as unresolvable.
func extractBasePathArg(call *sitter.Node, src []byte) (string, bool) {
	if call == nil || call.Type() != "function_call_expression" {
		return "", false
	}
	name := findChildByType(call, "name")
	if name == nil || nodeText(name, src) != "base_path" {
		return "", false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", false
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return "", false
	}
	if findChildByType(argNodes[0], "string") == nil {
		return "", false
	}
	return extractStringContent(argNodes[0], src), true
}
