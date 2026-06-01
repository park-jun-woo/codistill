//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what NestJS 프로젝트를 스캔하여 엔드포인트를 추출한다
package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a NestJS project root and extracts endpoints.
// Pass 1: collect routes from @Controller classes.
// Pass 2: resolve DTO types from import paths.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	tsFiles, err := findTSFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding ts files: %w", err)
	}
	if len(tsFiles) == 0 {
		fmt.Fprintf(os.Stderr, "warning: nestjs scanner found no .ts files under %s (checked %s and root itself)\n", absRoot, filepath.Join(absRoot, "src"))
		return &scanner.ScanResult{}, nil
	}
	globalPrefix := detectGlobalPrefix(absRoot)
	uriVersioning, defaultVersion := detectURIVersioning(absRoot)
	allControllers := collectControllers(tsFiles, absRoot)
	if len(allControllers) == 0 {
		fmt.Fprintf(os.Stderr, "warning: nestjs scanner found %d .ts files but no @Controller classes under %s\n", len(tsFiles), absRoot)
	}
	endpoints, dtoReqs := buildAllEndpoints(globalPrefix, uriVersioning, defaultVersion, allControllers, absRoot)
	schemas := resolveAllDTOs(dtoReqs, endpoints)
	return &scanner.ScanResult{Endpoints: endpoints, Schemas: schemas}, nil
}
