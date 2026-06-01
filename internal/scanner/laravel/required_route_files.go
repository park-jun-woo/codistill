//ff:func feature=scan type=extract control=iteration dimension=2 topic=laravel
//ff:what 라우트 파일 내부의 require/require_once로 로드되는 형제 라우트 파일들을 (감싼 그룹의 prefix/middleware와 함께) 수집한다
package laravel

// requiredRouteFiles scans a route file for require/require_once statements that
// pull in sibling route files (the pattern Bagisto's Admin/Shop route files use:
// a bare require 'auth-routes.php' at top level, plus require 'sales-routes.php'
// nested inside Route::group(['prefix'=>'admin'], fn)). Each include is resolved
// against this file's directory and returned with the prefix/middleware that
// applies at the require site: basePrefix/baseMW (what the file itself inherited)
// joined with every enclosing Route::group closure. Includes whose argument is
// not a static path (variable, etc.) are skipped.
func requiredRouteFiles(fi fileInfo, basePrefix string, baseMW []string) []routeFileRef {
	var refs []routeFileRef
	for _, t := range []string{"require_expression", "require_once_expression"} {
		for _, req := range findAllByType(fi.root, t) {
			rel, ok := requireIncludePath(req, fi.relPath, fi.src)
			if !ok {
				continue
			}
			prefix, mw := enclosingGroupModifiers(req, fi, basePrefix, baseMW)
			refs = append(refs, routeFileRef{relPath: rel, prefix: prefix, middleware: mw})
		}
	}
	return refs
}
