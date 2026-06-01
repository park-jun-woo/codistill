//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 데코레이터 패스: 파싱된 모든 파일에서 데코레이터 컨트롤러 라우트를 추출한다
package express

import (
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// scanDecoratorPass walks every parsed file looking for @RestController/
// @Controller-decorated classes with @Get/@Post/... method decorators
// (@n8n/decorators-style routing) and returns the synthesized endpoints. This
// is a separate pass from raw express route extraction; files with no such
// decorators contribute nothing.
func scanDecoratorPass(ctx *scanContext, absRoot string) []scanner.Endpoint {
	paths := make([]string, 0, len(ctx.parsed))
	for path := range ctx.parsed {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	var endpoints []scanner.Endpoint
	for _, path := range paths {
		fi := ctx.parsed[path]
		relPath := path
		if rel, err := filepath.Rel(absRoot, path); err == nil {
			relPath = rel
		}
		endpoints = append(endpoints, extractDecoratorEndpoints(fi, relPath)...)
	}
	return endpoints
}
