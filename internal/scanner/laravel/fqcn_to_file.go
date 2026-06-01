//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what FQCN과 composer psr-4 매핑으로 클래스 파일의 상대경로 후보를 산출한다
package laravel

import (
	"path"
	"strings"
)

// fqcnToFile maps a fully-qualified class name to its source file's relative
// path using the composer PSR-4 namespace->directory mappings. For a prefix
// "App\\" -> "app/" and FQCN "App\\Http\\Controllers\\Api\\UserController" it
// yields "app/Http/Controllers/Api/UserController.php". The longest matching
// prefix wins (so a more specific namespace mapping overrides a broader one).
// It returns ok=false when no PSR-4 prefix matches; callers then fall back to
// the hardcoded candidate paths.
func fqcnToFile(fqcn string, psr4 map[string]string) (string, bool) {
	fqcn = strings.TrimLeft(strings.TrimSpace(fqcn), "\\")
	if fqcn == "" {
		return "", false
	}
	bestPrefix, bestDir := "", ""
	for prefix, dir := range psr4 {
		np := strings.TrimLeft(prefix, "\\")
		if !strings.HasPrefix(fqcn, np) {
			continue
		}
		if len(np) >= len(bestPrefix) {
			bestPrefix, bestDir = np, dir
		}
	}
	if bestPrefix == "" {
		return "", false
	}
	rest := strings.TrimPrefix(fqcn, bestPrefix)
	relParts := strings.Split(rest, "\\")
	rel := path.Join(append([]string{strings.TrimRight(bestDir, "/")}, relParts...)...)
	return rel + ".php", true
}
