//ff:func feature=scan type=convert control=selection topic=quarkus
//ff:what 본문 컬렉션 제네릭(List<X> 등)의 내부 타입과 종류(none/array/map)를 판별한다
package quarkus

// normalizeBodyCollection classifies a Java collection generic used as a request
// body type. For List/Set/Collection/Iterable variants it returns the inner DTO
// type X and bodyCollectionArray so the caller can keep "X" for DTO resolution
// while marking the body as an array. For Map/HashMap it returns bodyCollectionMap
// so the caller can fall back to a free-form object. Otherwise it returns the
// input unchanged with bodyCollectionNone. This keeps the raw generic string
// (e.g. "List<CustomFieldJson>") out of OpenAPI schema keys.
func normalizeBodyCollection(typeName string) (inner string, kind bodyCollectionKind) {
	switch {
	case isBodyMapType(typeName):
		return "", bodyCollectionMap
	case isBodyListType(typeName):
		arg := extractGenericInner(typeName)
		if arg == "" {
			return "", bodyCollectionMap
		}
		return arg, bodyCollectionArray
	default:
		return typeName, bodyCollectionNone
	}
}
