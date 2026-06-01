//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what argument 노드의 첫 인자가 문자열 리터럴(빈 문자열 포함)인지 판별한다
package laravel

import sitter "github.com/smacker/go-tree-sitter"

// argIsStringLiteral reports whether the argument node holds a string literal
// (including an empty literal like '' or ""). Non-string first arguments such as
// variables or callables return false so callers can skip non-route calls.
func argIsStringLiteral(arg *sitter.Node) bool {
	for i := 0; i < int(arg.ChildCount()); i++ {
		switch arg.Child(i).Type() {
		case "string", "encapsed_string":
			return true
		}
	}
	return false
}
