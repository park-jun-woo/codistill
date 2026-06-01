//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 한 파일의 매크로 라우트 중 .service() 실등록된 핸들러만 엔드포인트로 변환해 추가한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func appendFileMacroEndpoints(endpoints []scanner.Endpoint, fi *fileInfo, sIdx structIndex, cache map[string][]scanner.Field, registeredHandlers map[string]bool) []scanner.Endpoint {
	for _, mr := range extractMacroRoutes(fi) {
		if !registeredHandlers[mr.handler] {
			continue
		}
		endpoints = append(endpoints, buildMacroEndpoint(mr, fi, sIdx, cache))
	}
	return endpoints
}
