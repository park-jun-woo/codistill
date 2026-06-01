//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 어노테이션 없는 POJO 파라미터를 request body로 감지한다
package quarkus

func classifyBodyParam(typeName, paramName string, ep *endpointInfo) {
	if typeName == "" || primitiveTypes[typeName] {
		return
	}
	if injectedParamTypes[typeName] {
		return
	}
	if ep.bodyType != "" {
		return
	}
	if inner, kind := normalizeBodyCollection(typeName); kind != bodyCollectionNone {
		// Map<...> → render a free-form object instead of a bogus "Map<...>"
		// schema key, so leave bodyType empty (buildRequest emits type:object).
		if kind == bodyCollectionMap {
			return
		}
		// List/Set/etc.<X> → keep the inner DTO type so DTO resolution and
		// schema lookup work on "X", and mark the body as an array; buildRequest
		// re-applies the slice marker for the shared array schema path.
		typeName = inner
		ep.bodyIsArray = true
	}
	ep.bodyType = typeName
	ep.bodyVarName = paramName
}
