//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 메서드 return 문들에서 응답 정보를 추출한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractResponsesFromMethod extracts response info from controller method return statements.
func extractResponsesFromMethod(absRoot string, cm *controllerMethod, parsedFiles map[string]*fileInfo) []scanner.Response {
	var responses []scanner.Response
	for _, retNode := range cm.returnNodes {
		resp := extractOneResponse(absRoot, retNode, cm.src, cm.srcFI, parsedFiles)
		if resp != nil {
			responses = append(responses, *resp)
		}
	}
	return appendAbortResponses(cm, responses)
}
