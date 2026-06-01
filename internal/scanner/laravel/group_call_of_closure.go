//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what group 클로저를 인자로 받는 바깥 호출 노드(member_call/scoped_call)를 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// groupCallOfClosure returns the call expression that takes closure as its
// group(...) argument — a member_call_expression for the chained
// Route::prefix()->group(fn) form, or a scoped_call_expression for the array
// Route::group([...], fn) form. It mirrors isGroupCallArgument's ascent but
// returns the call node so the caller can read its prefix/middleware.
func groupCallOfClosure(closure *sitter.Node) *sitter.Node {
	for a := closure.Parent(); a != nil; a = a.Parent() {
		switch a.Type() {
		case "argument", "arguments":
			continue
		case "member_call_expression", "scoped_call_expression":
			return a
		default:
			return nil
		}
	}
	return nil
}
