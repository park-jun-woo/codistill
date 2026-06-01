//ff:func feature=scan type=extract control=sequence
//ff:what operationId 선두의 비식별자 문자(`~`·`/` 등)를 제거하고 빈 경우 method 폴백한다
package scanner

import "strings"

func sanitizeOperationID(ep Endpoint, id string) string {
	trimmed := strings.TrimLeftFunc(id, func(r rune) bool {
		return !(r == '_' ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9'))
	})
	if trimmed == "" {
		// 전부 비식별자 문자면 빈 id 방지를 위해 method 기반으로 폴백한다.
		return methodPrefixedID(ep.Method, trimmed)
	}
	return trimmed
}
