//ff:func feature=scan type=parse control=iteration dimension=1 topic=fastapi
//ff:what Python 파일 목록을 모두 파싱하고 크로스파일 prefix를 병합한다
package fastapi

import (
	"os"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAllFiles parses all Python files and returns fileInfo list.
// A pre-pass first collects APIRouter subclass names across all files so that
// custom routers (e.g. UserAPIRouter(APIRouter)) defined in one file but used
// in another are recognized. After the per-file parse, it runs a cross-file
// prefix merge pass to resolve include_router chains spanning multiple files.
func parseAllFiles(absRoot string, pyFiles []string) []fileInfo {
	roots := make([]*sitter.Node, len(pyFiles))
	srcs := make([][]byte, len(pyFiles))
	for i, f := range pyFiles {
		src, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		root, err := parsePython(src)
		if err != nil {
			continue
		}
		roots[i] = root
		srcs[i] = src
	}
	routerSubclassNames := collectRouterSubclasses(roots, srcs)

	var files []fileInfo
	for _, f := range pyFiles {
		fi, err := parseFile(absRoot, f, routerSubclassNames)
		if err != nil {
			continue
		}
		files = append(files, *fi)
	}
	mergeCrossFilePrefixes(absRoot, files)
	return files
}
