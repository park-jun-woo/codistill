//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what add_resource/튜플 등록과 Resource 카탈로그를 결합해 라우트를 추출한다
package flask

// extractResourceRoutes resolves Flask-RESTful class-based routes. It collects
// add_resource(Resource, path...) calls and configure_api_from_blueprint-style
// tuple-list registrations across all files, looks each Resource class up in the
// catalog, and emits one routeInfo per (registered path × HTTP method def). The
// blueprint prefix is composed via bpPrefixes when a registration carries a
// blueprint variable; the handler label is "ClassName.METHOD".
func extractResourceRoutes(files []fileInfo, catalog resourceClassCatalog, bpPrefixes blueprintPrefix, nsPrefixes namespacePrefix) []routeInfo {
	var routes []routeInfo
	for _, fi := range files {
		regs := collectAddResource(fi.root, fi.src)
		regs = append(regs, collectConfigureAPITuples(fi.root, fi.src)...)
		routes = append(routes, regsToRouteInfo(regs, catalog, bpPrefixes, nsPrefixes)...)
	}
	return routes
}
