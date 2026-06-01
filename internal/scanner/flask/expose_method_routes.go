//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 단일 @expose path/methods를 base_url과 합성해 routeInfo로 변환한다
package flask

// exposeMethodRoutes composes one @expose path against an API class's base_url
// (combinePath -> OpenAPI form, URL params extracted) and returns one routeInfo
// per HTTP method, labeled "ClassName.method". Used by exposeRoutesFromMethod.
func exposeMethodRoutes(api appbuilderAPIInfo, methodName, rawPath string, methods []string, line int) []routeInfo {
	fullPath := combinePath(api.baseURL, rawPath)
	openAPIPath := flaskPathToOpenAPI(fullPath)
	params := extractURLParams(fullPath)
	var routes []routeInfo
	for _, m := range methods {
		routes = append(routes, routeInfo{
			method:  m,
			path:    openAPIPath,
			handler: api.name + "." + methodName,
			file:    api.file,
			line:    line,
			params:  params,
		})
	}
	return routes
}
