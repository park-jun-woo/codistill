//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일기반 라우팅 패스: src/api/**/route.ts 파일에서 export const VERB 라우트를 추출한다
package express

import (
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// scanFilebasedPass walks every parsed file, keeps only Medusa-v2 file-based
// routing files (src/api/**/route.{ts,js}), and synthesizes an endpoint per
// `export const <VERB>` handler. This is a separate pass from raw express
// route extraction; ordinary express files contribute nothing because the
// isMedusaRouteFile guard rejects them.
func scanFilebasedPass(ctx *scanContext, absRoot string) []scanner.Endpoint {
	paths := make([]string, 0, len(ctx.parsed))
	for path := range ctx.parsed {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	var endpoints []scanner.Endpoint
	for _, path := range paths {
		relPath := path
		if rel, err := filepath.Rel(absRoot, path); err == nil {
			relPath = rel
		}
		relPath = filepath.ToSlash(relPath)
		if !isMedusaRouteFile(relPath) {
			continue
		}
		endpoints = append(endpoints, extractFilebasedEndpoints(ctx.parsed[path], relPath)...)
	}
	return endpoints
}
