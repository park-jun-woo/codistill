//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 배열형 Route::group([...], fn) scoped 호출의 옵션 배열에서 prefix/middleware를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// scopedGroupModifiers reads the prefix/middleware from the options array of an
// array-form Route::group(['prefix'=>'x', ...], fn) scoped_call_expression. It
// returns empty values when the call is not a group call or has no array option.
func scopedGroupModifiers(call *sitter.Node, fi fileInfo) (string, []string) {
	if secondScopedName(call, fi.src) != "group" {
		return "", nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", nil
	}
	for _, arg := range childrenOfType(args, "argument") {
		if arr := findChildByType(arg, "array_creation_expression"); arr != nil {
			return extractGroupArrayModifier(arr, fi)
		}
	}
	return "", nil
}
