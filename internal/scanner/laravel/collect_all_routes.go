//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 라우트 파일(Provider 등록 분할 파일 + routes/api.php·web.php 폴백)에서 라우트를 모은다
package laravel

import "strings"

// collectAllRoutes gathers routes from every route file. Files loaded by a
// RouteServiceProvider via ->group(base_path('routes/X.php')) or
// require base_path('routes/X.php') are discovered dynamically and collected
// with the provider's prefix/middleware, which is the source of truth for the
// applied prefix. The conventional routes/api.php and routes/web.php act as a
// fallback for apps without explicit provider loads; any file already
// registered by the provider is not collected twice. The filename heuristic
// (api.php -> /api) is applied only when the provider scan found no route file
// registrations at all, so apps whose provider loads api.php without a prefix
// (e.g. ->middleware('api')->group(...)) no longer get a spurious /api prefix.
func collectAllRoutes(parsedFiles map[string]*fileInfo) []routeInfo {
	var routes []routeInfo
	collected := make(map[string]bool)

	providerRefs := providerRouteFiles(parsedFiles)
	for _, ref := range providerRefs {
		fi, ok := parsedFiles[ref.relPath]
		if !ok || collected[ref.relPath] {
			continue
		}
		collected[ref.relPath] = true
		routes = append(routes, collectRoutes(*fi, ref.prefix, ref.middleware)...)
		routes = append(routes, collectAPIResource(*fi, ref.prefix, ref.middleware)...)
		routes = append(routes, extractRouteGroups(*fi, ref.prefix, ref.middleware)...)
	}

	providerLoaded := len(providerRefs) > 0
	for _, rf := range []string{"routes/api.php", "routes/web.php"} {
		fi, ok := parsedFiles[rf]
		if !ok || collected[rf] {
			continue
		}
		collected[rf] = true
		prefix := ""
		if !providerLoaded && strings.HasSuffix(rf, "api.php") {
			prefix = "api"
		}
		routes = append(routes, collectRoutes(*fi, prefix, nil)...)
		routes = append(routes, collectAPIResource(*fi, prefix, nil)...)
		routes = append(routes, extractRouteGroups(*fi, prefix, nil)...)
	}

	return routes
}
