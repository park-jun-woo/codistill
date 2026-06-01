//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what relPath가 app/Providers/ 아래의 Provider PHP 파일인지 판정한다
package laravel

import "strings"

// isRouteServiceProvider reports whether relPath is a provider PHP file under
// app/Providers/, the conventional home of Laravel's RouteServiceProvider and
// any custom route-registering providers (e.g. Akaunting's app/Providers/Route.php).
func isRouteServiceProvider(relPath string) bool {
	p := strings.ReplaceAll(relPath, "\\", "/")
	return strings.HasPrefix(p, "app/Providers/") && strings.HasSuffix(p, ".php")
}
