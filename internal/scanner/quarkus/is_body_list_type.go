//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 본문 타입이 List/Set/Collection 계열 제네릭(array로 표현)인지 확인한다
package quarkus

import "strings"

func isBodyListType(typeName string) bool {
	for _, prefix := range []string{
		"List<", "Set<", "Collection<", "Iterable<",
		"ArrayList<", "LinkedList<", "HashSet<", "LinkedHashSet<", "SortedSet<", "TreeSet<",
	} {
		if strings.HasPrefix(typeName, prefix) && strings.HasSuffix(typeName, ">") {
			return true
		}
	}
	return false
}
