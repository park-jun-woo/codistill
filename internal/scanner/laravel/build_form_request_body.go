//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what FormRequest 참조에서 필드를 추출해 요청 본문(body)을 구성한다(필드 없으면 nil)
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildFormRequestBody(absRoot string, cm *controllerMethod, parsedFiles map[string]*fileInfo) *scanner.Body {
	fields := extractFormRequest(absRoot, cm.formRequestRef, cm.srcFI, parsedFiles)
	if len(fields) == 0 {
		return nil
	}
	return &scanner.Body{
		VarName:  "request",
		Method:   "json",
		TypeName: cm.formRequestRef,
		Fields:   fields,
	}
}
