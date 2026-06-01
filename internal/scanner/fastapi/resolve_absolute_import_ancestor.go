//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 모듈 첫 세그먼트가 조상 디렉터리에 존재하는 패키지 루트를 거슬러 찾아 import를 해석한다
package fastapi

import (
	"os"
	"path/filepath"
	"strings"
)

// resolveAbsoluteImportAncestor resolves an absolute Python import path by
// walking up from absRoot to find an ancestor directory that contains the
// module's first segment as a sub-directory, then treating that ancestor as the
// package root. This handles subpath scans where the scan root differs from the
// import (package) root, e.g. scanning "backend/open_webui" while imports read
// "open_webui.routers". Returns "" if no matching ancestor is found.
func resolveAbsoluteImportAncestor(absRoot, module string) string {
	if module == "" || strings.HasPrefix(module, ".") {
		return ""
	}
	segs := strings.Split(module, ".")
	firstSeg := segs[0]
	relPath := strings.ReplaceAll(module, ".", string(filepath.Separator))

	// 무한 상향 방지: 모듈 세그먼트 수 + 4 만큼만 거슬러 올라간다.
	maxUp := len(segs) + 4
	dir := absRoot
	for i := 0; i < maxUp; i++ {
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
		if _, err := os.Stat(filepath.Join(dir, firstSeg)); err != nil {
			continue
		}
		fullPath := filepath.Join(dir, relPath)
		pyFile := fullPath + ".py"
		if _, err := os.Stat(pyFile); err == nil {
			return pyFile
		}
		initFile := filepath.Join(fullPath, "__init__.py")
		if _, err := os.Stat(initFile); err == nil {
			return initFile
		}
	}
	return ""
}
