//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 라우트 파일(Provider 등록 분할 파일 + 내부 require 형제 파일 + routes/api.php·web.php 폴백)에서 라우트를 모은다
package laravel

import "strings"

// collectAllRoutes gathers routes from every route file. Files loaded by a
// RouteServiceProvider via ->group(base_path('routes/X.php')), $this->
// loadRoutesFrom(__DIR__.'/../Routes/X.php'), or require base_path(...) are
// discovered dynamically and collected with the provider's prefix/middleware.
// Each collected route file is then scanned for require/require_once of sibling
// route files (Bagisto's Admin/Shop split-by-require convention), which are
// followed transitively with the enclosing group's prefix/middleware. The
// conventional routes/api.php and routes/web.php act as a fallback for apps
// without explicit provider loads; any file already collected is not collected
// twice. The filename heuristic (api.php -> /api) is applied only when the
// provider scan found no route file registrations at all.
func collectAllRoutes(parsedFiles map[string]*fileInfo) []routeInfo {
	var routes []routeInfo
	collected := make(map[string]bool)

	collect := func(ref routeFileRef) {
		queue := []routeFileRef{ref}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			fi, ok := parsedFiles[cur.relPath]
			if !ok || collected[cur.relPath] {
				continue
			}
			collected[cur.relPath] = true
			routes = append(routes, collectRoutes(*fi, cur.prefix, cur.middleware)...)
			routes = append(routes, collectAPIResource(*fi, cur.prefix, cur.middleware)...)
			routes = append(routes, extractRouteGroups(*fi, cur.prefix, cur.middleware)...)
			queue = append(queue, requiredRouteFiles(*fi, cur.prefix, cur.middleware)...)
		}
	}

	providerRefs := providerRouteFiles(parsedFiles)
	for _, ref := range providerRefs {
		collect(ref)
	}

	providerLoaded := len(providerRefs) > 0
	for _, rf := range []string{"routes/api.php", "routes/web.php"} {
		if collected[rf] {
			continue
		}
		if _, ok := parsedFiles[rf]; !ok {
			continue
		}
		prefix := ""
		if !providerLoaded && strings.HasSuffix(rf, "api.php") {
			prefix = "api"
		}
		collect(routeFileRef{relPath: rf, prefix: prefix})
	}

	return routes
}
