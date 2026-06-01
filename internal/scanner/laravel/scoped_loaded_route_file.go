//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::group(base_path('routes/X.php')) scoped_call에서 라우트 파일 참조를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// scopedLoadedRouteFile inspects a scoped_call_expression of the form
// Route::group(base_path('routes/X.php')) — a group call with no chained
// ->prefix()/->middleware() — and returns the loaded route file. Such a bare
// group carries no prefix/middleware (chained forms are handled via member calls).
func scopedLoadedRouteFile(call *sitter.Node, fi fileInfo) (routeFileRef, bool) {
	if secondScopedName(call, fi.src) != "group" {
		return routeFileRef{}, false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return routeFileRef{}, false
	}
	rel, ok := basePathArgFromArguments(args, fi.src)
	if !ok {
		rel, ok = dirRelativeRouteFile(args, fi.relPath, fi.src)
	}
	if !ok {
		return routeFileRef{}, false
	}
	return routeFileRef{relPath: rel}, true
}
