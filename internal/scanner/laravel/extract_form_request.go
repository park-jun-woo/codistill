//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what FormRequest::rules() 배열에서 필드명과 유효성 규칙을 추출한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractFormRequest finds a FormRequest class and extracts fields from its rules() method.
func extractFormRequest(absRoot, className string, srcFI *fileInfo, parsedFiles map[string]*fileInfo) []scanner.Field {
	fi := findFormRequestFile(absRoot, className, srcFI, parsedFiles)
	if fi == nil {
		return nil
	}
	return extractRulesFromFile(fi, className)
}
