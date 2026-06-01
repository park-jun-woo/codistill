//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 클래스 HTTP 메서드 목록을 routeInfo 슬라이스로 변환한다
package flask

// classRoutesToRouteInfo converts a class's HTTP method defs into routeInfo
// entries. The route path is composed as prefix + path (combinePath), then
// normalized to OpenAPI format (flaskPathToOpenAPI) with URL params extracted.
// handlerPrefix is prepended to each method name to form the handler label
// (e.g. "UserAPI." + "get"). Used by Phase170/171/172 class-based scanners.
func classRoutesToRouteInfo(methods []classMethod, path, prefix, handlerPrefix, file string) []routeInfo {
	fullPath := combinePath(prefix, path)
	openAPIPath := flaskPathToOpenAPI(fullPath)
	params := extractURLParams(fullPath)

	var routes []routeInfo
	for _, m := range methods {
		routes = append(routes, routeInfo{
			method:  m.name,
			path:    openAPIPath,
			handler: handlerPrefix + m.name,
			file:    file,
			line:    m.line,
			params:  params,
		})
	}
	return routes
}
