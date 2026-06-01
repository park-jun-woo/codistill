//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what require/require_once base_path('routes/X.php') 식에서 라우트 파일 참조를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// requireLoadedRouteFile inspects a require_expression / require_once_expression
// and, when it includes base_path('routes/X.php'), returns the loaded route file
// with empty prefix/middleware (a bare require carries no group modifiers).
func requireLoadedRouteFile(node *sitter.Node, src []byte) (routeFileRef, bool) {
	call := findChildByType(node, "function_call_expression")
	if call == nil {
		return routeFileRef{}, false
	}
	rel, ok := extractBasePathArg(call, src)
	if !ok {
		return routeFileRef{}, false
	}
	return routeFileRef{relPath: rel}, true
}
