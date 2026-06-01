//ff:func feature=scan type=extract control=iteration dimension=2 topic=laravel
//ff:what app/Providers/ 및 패키지 *ServiceProvider.php에서 base_path/__DIR__ 상대경로로 로드되는 라우트 파일 목록을 수집한다
package laravel

// providerRouteFiles scans RouteServiceProvider-style files under app/Providers/
// and modular package providers (*ServiceProvider.php, e.g. under packages/**/
// src/Providers/) for split route files loaded via
// Route::...->group(base_path('routes/X.php')), Route::...->group(__DIR__ .
// '/../Routes/X.php'), $this->loadRoutesFrom(__DIR__ . '/../Routes/X.php'), or
// require/require_once base_path('routes/X.php'), returning each loaded file
// with the prefix/middleware the provider applies.
// Refs are deduped by relPath (first occurrence wins) so a file loaded twice is
// collected once.
func providerRouteFiles(parsedFiles map[string]*fileInfo) []routeFileRef {
	var refs []routeFileRef
	seen := make(map[string]bool)
	add := func(ref routeFileRef, ok bool) {
		if !ok || seen[ref.relPath] {
			return
		}
		seen[ref.relPath] = true
		refs = append(refs, ref)
	}
	for relPath, fi := range parsedFiles {
		if !isRouteServiceProvider(relPath) && !isPackageServiceProvider(relPath) {
			continue
		}
		for _, mc := range findAllByType(fi.root, "member_call_expression") {
			ref, ok := groupLoadedRouteFile(mc, *fi)
			add(ref, ok)
			lref, lok := loadRoutesFromFile(mc, *fi)
			add(lref, lok)
		}
		for _, sc := range findAllByType(fi.root, "scoped_call_expression") {
			ref, ok := scopedLoadedRouteFile(sc, *fi)
			add(ref, ok)
		}
		for _, req := range findAllByType(fi.root, "require_expression") {
			ref, ok := requireLoadedRouteFile(req, fi.src)
			add(ref, ok)
		}
		for _, req := range findAllByType(fi.root, "require_once_expression") {
			ref, ok := requireLoadedRouteFile(req, fi.src)
			add(ref, ok)
		}
	}
	return refs
}
