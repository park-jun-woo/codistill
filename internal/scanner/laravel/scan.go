//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Laravel 프로젝트를 3-pass로 스캔하여 엔드포인트를 추출한다
package laravel

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a Laravel project root and extracts endpoints.
// Pass 1: collect route structure from routes/api.php and routes/web.php.
// Pass 2: resolve controller methods — extract parameter types and return info.
// Pass 3: resolve FormRequest rules and Resource response fields.
//
// Known limitation: dynamic route macros such as Auth::routes() are not
// expanded. They register routes at runtime via framework internals, which is
// outside the reach of static analysis, so the routes they emit are omitted.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}

	// Stage 1: parse only route-source files (routes/**, providers,
	// *ServiceProvider.php, **/Routes/**). Controllers/FormRequests/Resources
	// are parsed lazily on demand in stage 2 (buildEndpoints), so unrelated PHP
	// files in large apps are never parsed up front.
	routeSrcFiles, err := findRouteSourceFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding route source files: %w", err)
	}
	if len(routeSrcFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	parsedFiles := parseAllPHPFiles(absRoot, routeSrcFiles)
	if len(parsedFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	routes := collectAllRoutes(parsedFiles)
	if len(routes) == 0 {
		return &scanner.ScanResult{}, nil
	}

	endpoints := buildEndpoints(absRoot, routes, parsedFiles)
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
