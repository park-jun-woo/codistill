//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 단일 등록의 각 path를 Resource 메서드와 결합해 routeInfo로 변환한다
package flask

// regPathsToRouteInfo expands one resource registration's paths into routeInfo
// entries, one per (path × HTTP method def). The handler label is "ClassName.METHOD".
func regPathsToRouteInfo(reg addResourceReg, info resourceClassInfo, prefix string) []routeInfo {
	var routes []routeInfo
	for _, path := range reg.paths {
		rs := classRoutesToRouteInfo(info.methods, path, prefix, reg.className+".", info.file)
		routes = append(routes, rs...)
	}
	return routes
}
