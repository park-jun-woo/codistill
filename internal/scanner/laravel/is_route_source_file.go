//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what relPath가 라우트 추출에 필요한 소스 파일(routes/**, app/Providers/**, *ServiceProvider.php, **/Routes/**, **/Http/routes.php)인지 판정한다
package laravel

import "strings"

// isRouteSourceFile reports whether relPath is a PHP file that may define or
// load routes and therefore must be parsed in stage 1. The set is intentionally
// narrow — route files (routes/**), service providers (app/Providers/** and any
// *ServiceProvider.php), and package route directories (**/Routes/**) — so that
// the thousands of unrelated PHP files in a large app (models, migrations,
// seeders, factories, Blade-ish PHP) are not parsed up front.
func isRouteSourceFile(relPath string) bool {
	p := strings.ReplaceAll(relPath, "\\", "/")
	switch {
	case strings.HasPrefix(p, "routes/"):
		return true
	case strings.HasPrefix(p, "app/Providers/"):
		return true
	case strings.HasSuffix(p, "ServiceProvider.php"):
		return true
	case strings.Contains(p, "/Routes/"):
		return true
	case strings.HasSuffix(p, "/Http/routes.php"):
		return true
	default:
		return false
	}
}
