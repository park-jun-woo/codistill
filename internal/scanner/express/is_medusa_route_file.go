//ff:func feature=scan type=extract control=sequence topic=express
//ff:what Medusa v2 파일기반 라우팅 대상(src/api/**/route.{ts,js}) 파일인지 판정한다
package express

import (
	"path/filepath"
	"strings"
)

// isMedusaRouteFile reports whether relPath is a Medusa-v2 file-based routing
// file: a file named route.ts/route.js located under an "api" directory. The
// guard prevents false positives in ordinary express repos: only files whose
// path contains an "api" segment and whose basename is route.{ts,js} qualify.
// relPath is expected to use forward slashes (filepath.Rel on the scan root).
func isMedusaRouteFile(relPath string) bool {
	base := filepath.Base(relPath)
	if base != "route.ts" && base != "route.js" {
		return false
	}
	segs := strings.Split(filepath.ToSlash(relPath), "/")
	// require an "api" directory somewhere before the file itself.
	return pathHasAPIDir(segs[:len(segs)-1])
}
