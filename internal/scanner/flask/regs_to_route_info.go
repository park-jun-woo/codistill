//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 등록 정보 목록을 카탈로그/prefix와 결합해 routeInfo 슬라이스로 변환한다
package flask

// regsToRouteInfo resolves each addResourceReg against the Resource catalog and
// emits routeInfo entries. Registrations whose class is absent from the catalog
// or has no HTTP method defs are skipped. The prefix is resolved per registration
// from the receiver variable: a blueprint prefix takes precedence, otherwise a
// flask_restx namespace prefix (resolveRegPrefix).
func regsToRouteInfo(regs []addResourceReg, catalog resourceClassCatalog, bpPrefixes blueprintPrefix, nsPrefixes namespacePrefix) []routeInfo {
	var routes []routeInfo
	for _, reg := range regs {
		info, ok := catalog[reg.className]
		if !ok || len(info.methods) == 0 {
			continue
		}
		prefix := resolveRegPrefix(reg.blueprintVar, bpPrefixes, nsPrefixes)
		routes = append(routes, regPathsToRouteInfo(reg, info, prefix)...)
	}
	return routes
}
