//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 본문 타입이 Map 계열 제네릭(free-form object로 폴백)인지 확인한다
package quarkus

import "strings"

func isBodyMapType(typeName string) bool {
	return strings.HasSuffix(typeName, ">") &&
		(strings.HasPrefix(typeName, "Map<") || strings.HasPrefix(typeName, "HashMap<"))
}
