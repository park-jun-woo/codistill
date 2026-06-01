//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 노드를 감싸는 모든 Route::group 클로저의 prefix/middleware를 바깥→안쪽 순으로 누적해 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// enclosingGroupModifiers walks upward from node and accumulates the
// prefix/middleware of every Route::group closure that wraps it, so a require
// statement buried inside Route::group(['prefix'=>'admin'], fn) inherits that
// group's prefix. Both the chained Route::prefix()->group(fn) and array
// Route::group([...], fn) forms are recognized. base prefix/middleware seed the
// result (the prefix/middleware applied to the requiring file itself), and
// nearer-enclosing groups are joined as inner segments after farther ones.
func enclosingGroupModifiers(node *sitter.Node, fi fileInfo, basePrefix string, baseMW []string) (string, []string) {
	var prefixes []string
	var mws [][]string
	for n := node.Parent(); n != nil; n = n.Parent() {
		if n.Type() != "anonymous_function_creation_expression" && n.Type() != "arrow_function" {
			continue
		}
		if !isGroupCallArgument(n, fi) {
			continue
		}
		call := groupCallOfClosure(n)
		if call == nil {
			continue
		}
		p, m := groupCallModifiers(call, fi)
		prefixes = append(prefixes, p)
		mws = append(mws, m)
	}
	prefix := basePrefix
	mw := baseMW
	for i := len(prefixes) - 1; i >= 0; i-- {
		prefix = joinGroupPrefix(prefix, prefixes[i])
		mw = mergeMiddleware(mw, mws[i])
	}
	return prefix, mw
}
