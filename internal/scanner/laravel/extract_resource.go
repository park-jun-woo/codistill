//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what JsonResource::toArray() 에서 응답 필드 키를 추출한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractResourceFields finds a JsonResource class and extracts field names from toArray().
func extractResourceFields(absRoot, className string, srcFI *fileInfo, parsedFiles map[string]*fileInfo) []scanner.Field {
	fi := findResourceFile(absRoot, className, srcFI, parsedFiles)
	if fi == nil {
		return nil
	}
	return extractToArrayFields(fi, className)
}
