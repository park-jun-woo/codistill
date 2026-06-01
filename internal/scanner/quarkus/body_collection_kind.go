//ff:type feature=scan type=model topic=quarkus
//ff:what 본문 컬렉션 제네릭의 분류 종류(none/array/map)
package quarkus

type bodyCollectionKind int

const (
	bodyCollectionNone bodyCollectionKind = iota
	bodyCollectionArray
	bodyCollectionMap
)
