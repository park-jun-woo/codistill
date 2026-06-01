//ff:func feature=scan type=extract control=sequence
//ff:what JoinPath 좌측 fold 1스텝(첫 인자는 그대로, 이후는 JoinPath)
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// foldPart returns the first part as-is, otherwise scanner.JoinPath(acc, part).
func foldPart(acc, part string, i int) string {
	if i == 0 {
		return part
	}
	return scanner.JoinPath(acc, part)
}
