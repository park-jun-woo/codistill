//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 체인형 Route::prefix('x')->middleware([..])->group(fn) member 호출 체인에서 prefix/middleware를 누적한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// chainedGroupModifiers walks the call chain of a chained-form
// Route::prefix('x')->middleware([...])->group(fn) member_call_expression and
// returns the accumulated prefix/middleware. It returns empty values when the
// outermost call is not a group call.
func chainedGroupModifiers(call *sitter.Node, fi fileInfo) (string, []string) {
	if lastMemberCallName(call, fi.src) != "group" {
		return "", nil
	}
	var prefix string
	var mw []string
	if inner := findChildByType(call, "scoped_call_expression"); inner != nil {
		walkChain(inner, fi, &prefix, &mw)
	}
	if inner := findChildByType(call, "member_call_expression"); inner != nil {
		walkChain(inner, fi, &prefix, &mw)
	}
	return prefix, mw
}
