//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what ->group(base_path('routes/X.php') | __DIR__ . '/../Routes/X.php') member_call에서 라우트 파일 참조(prefix/middleware 포함)를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// groupLoadedRouteFile inspects a member_call_expression of the form
// Route::prefix('p')->middleware([...])->group(base_path('routes/X.php')) and,
// when its ->group(...) argument is a base_path() string or a __DIR__-relative
// concatenation (resolved against the provider file's directory), returns the
// loaded route file together with the prefix/middleware accumulated along the
// chain. The __DIR__ form is the convention modular packages use to load their
// own Routes/ files (e.g. Bagisto's package service providers).
func groupLoadedRouteFile(mc *sitter.Node, fi fileInfo) (routeFileRef, bool) {
	if lastMemberCallName(mc, fi.src) != "group" {
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
	prefix := ""
	var mw []string
	if inner := findChildByType(mc, "scoped_call_expression"); inner != nil {
		walkChain(inner, fi, &prefix, &mw)
	}
	if inner := findChildByType(mc, "member_call_expression"); inner != nil {
		walkChain(inner, fi, &prefix, &mw)
	}
	return routeFileRef{relPath: rel, prefix: prefix, middleware: mw}, true
}
