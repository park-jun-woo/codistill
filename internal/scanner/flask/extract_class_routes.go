//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what flask_restx @ns.route 클래스 데코레이터 라우트를 전 파일에서 추출한다
package flask

// extractClassRoutes resolves flask_restx class-based routes declared with
// @ns.route("/path") decorators on a Resource subclass. For each decorated
// class_definition whose superclass list contains Resource, every *.route
// decorator yields one path; the HTTP methods come from the class body
// (classHTTPMethods). The namespace variable in the decorator (e.g. "inner_api_ns"
// in @inner_api_ns.route) is resolved against nsPrefixes to compose
// namespace_prefix + route_path; an unknown namespace falls back to no prefix.
// The handler label is "ClassName.METHOD".
func extractClassRoutes(files []fileInfo, nsPrefixes namespacePrefix) []routeInfo {
	var routes []routeInfo
	for _, fi := range files {
		aliases := collectImportAliases(fi.root, fi.src)
		for _, dd := range findAllByType(fi.root, "decorated_definition") {
			routes = append(routes, classRoutesFromDecorated(dd, fi, aliases, nsPrefixes)...)
		}
	}
	return routes
}
