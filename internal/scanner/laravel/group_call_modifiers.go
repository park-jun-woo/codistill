//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what group 호출 노드에서 prefix/middleware를 추출한다 (배열형 Route::group([..],fn) + 체인형 Route::prefix()->group(fn))
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// groupCallModifiers extracts the prefix and middleware a Route::group call
// applies, dispatching on its form: the array form Route::group([...], fn)
// (a scoped_call_expression) reads the options array, while the chained form
// Route::prefix('x')->middleware([...])->group(fn) (a member_call_expression)
// walks the call chain. It returns empty values for any non-group call.
func groupCallModifiers(call *sitter.Node, fi fileInfo) (string, []string) {
	switch call.Type() {
	case "scoped_call_expression":
		return scopedGroupModifiers(call, fi)
	case "member_call_expression":
		return chainedGroupModifiers(call, fi)
	}
	return "", nil
}
