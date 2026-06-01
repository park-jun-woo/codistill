//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 경로에서 첫 번째 의미 있는 segment를 추출한다
package scanner

import (
	"strings"
)

func firstPathSegment(path string) string {
	segments := strings.Split(path, "/")
	for _, seg := range segments {
		if seg == "" || seg == "api" {
			continue
		}
		if strings.HasPrefix(seg, "v") && len(seg) <= 3 {
			continue
		}
		if strings.HasPrefix(seg, "{") || strings.HasPrefix(seg, ":") {
			continue
		}
		// Skip unresolved template-literal fragments (`${PREFIX}`) so they do
		// not pollute the operationId prefix. See Phase145.
		if strings.HasPrefix(seg, "$") {
			continue
		}
		return seg
	}
	return ""
}
