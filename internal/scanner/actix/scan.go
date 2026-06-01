//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what Actix-web 프로젝트를 스캔하여 엔드포인트를 추출한다
package actix

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	rsFiles, err := findRsFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding rs files: %w", err)
	}
	if len(rsFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, rsFiles)
	if len(files) == 0 {
		return &scanner.ScanResult{}, nil
	}

	sIdx := buildStructIndex(files)
	fieldCache := make(map[string][]scanner.Field)
	handlerFuncs := make(map[string]*handlerInfo)

	registeredHandlers := collectRegisteredHandlers(files)

	endpoints := scanMacroEndpoints(files, sIdx, fieldCache, handlerFuncs, registeredHandlers)
	endpoints = append(endpoints, scanBuilderEndpoints(files, sIdx, fieldCache, handlerFuncs)...)
	applyScopePrefixes(files, endpoints)

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
