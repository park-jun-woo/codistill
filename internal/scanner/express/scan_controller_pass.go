//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 커스텀 컨트롤러 패스: 파싱된 모든 파일에서 Controller 서브클래스의 this 기반 라우트를 추출한다
package express

import (
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// scanControllerPass walks every parsed file looking for classes that extend a
// base Controller (Unleash-style) and registers routes via this.route({...}) or
// this.<method>("..."). It is a separate pass from raw express route extraction
// and the decorator pass; files with no Controller subclass contribute nothing.
func scanControllerPass(ctx *scanContext, absRoot string) []scanner.Endpoint {
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
		endpoints = append(endpoints, extractControllerEndpoints(fi, relPath)...)
	}
	return endpoints
}
