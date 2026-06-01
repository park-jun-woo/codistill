//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what $this->loadRoutesFrom(__DIR__ . '/../Routes/X.php') member_call에서 라우트 파일 참조를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// loadRoutesFromFile inspects a member_call_expression of the form
// $this->loadRoutesFrom(__DIR__ . '/../Routes/X.php') — the convention modular
// package service providers (e.g. Bagisto's payment gateways) use to register
// their own routes — and returns the loaded route file resolved against the
// provider file's directory. base_path('routes/X.php') arguments are also
// accepted. A loadRoutesFrom call carries no prefix/middleware of its own;
// any group wrapping happens inside the loaded file.
func loadRoutesFromFile(mc *sitter.Node, fi fileInfo) (routeFileRef, bool) {
	if memberCallName(mc, fi.src) != "loadRoutesFrom" {
		return routeFileRef{}, false
	}
	args := findChildByType(mc, "arguments")
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
