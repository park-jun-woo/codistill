//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what 단일 Python 파일을 파싱하여 fileInfo를 반환한다
package fastapi

import (
	"os"
	"path/filepath"
)

// parseFile parses a single Python file and returns its fileInfo.
// routerSubclassNames is the repo-wide set of APIRouter subclass names used to
// recognize custom router instantiations across files.
func parseFile(absRoot, absPath string, routerSubclassNames map[string]bool) (*fileInfo, error) {
	src, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	root, err := parsePython(src)
	if err != nil {
		return nil, err
	}
	relPath, _ := filepath.Rel(absRoot, absPath)
	return &fileInfo{
		absPath:    absPath,
		relPath:    relPath,
		src:        src,
		root:       root,
		imports:    extractImports(root, src),
		prefixes:   resolveRouterPrefixes(root, src, routerSubclassNames),
		routerDeps: resolveRouterDeps(root, src, routerSubclassNames),
		hidden:     resolveHiddenRouters(root, src, routerSubclassNames),
		models:     findAllPydanticModels(root, src),
	}, nil
}
