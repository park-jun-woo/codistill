//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 클래스명에서 파일 경로를 추적하고 메서드를 파싱한다 (캐시→PSR-4→하드코딩 폴백)
package laravel

import (
	"path/filepath"
)

// resolveController finds and parses a controller file from its class name.
// Resolution order: (1) the parse cache (route-stage files plus any class
// already lazily parsed), (2) PSR-4 — the route file's `use` import gives the
// FQCN which composer's psr-4 map turns into an exact path (recovers
// non-standard locations the old full parse covered), (3) the conventional
// app/Http/Controllers/** hardcoded candidates. srcFI is the route file naming
// the controller; it may be nil (PSR-4 step is then skipped).
func resolveController(absRoot, className string, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *fileInfo {
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}

	if fi := resolveClassViaPSR4(absRoot, className, srcFI, parsedFiles); fi != nil {
		return fi
	}

	candidates := []string{
		filepath.Join(absRoot, "app", "Http", "Controllers", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "Api", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "API", className+".php"),
	}
	for _, candidate := range candidates {
		if fi := parseControllerCandidate(absRoot, candidate); fi != nil {
			return fi
		}
	}
	return nil
}
