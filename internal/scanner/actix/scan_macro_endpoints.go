//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Pass 1: .service() 실등록된 매크로 라우트만 엔드포인트로 추출하고 핸들러 인덱스를 채운다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanMacroEndpoints(files []*fileInfo, sIdx structIndex, cache map[string][]scanner.Field, handlerFuncs map[string]*handlerInfo, registeredHandlers map[string]bool) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, fi := range files {
		endpoints = appendFileMacroEndpoints(endpoints, fi, sIdx, cache, registeredHandlers)
		collectHandlerFuncs(fi, handlerFuncs)
	}
	return endpoints
}
