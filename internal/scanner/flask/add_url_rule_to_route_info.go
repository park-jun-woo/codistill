//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 단일 add_url_rule 등록을 prefix와 결합해 메서드별 routeInfo로 변환한다
package flask

// addURLRuleToRouteInfo resolves one add_url_rule registration into routeInfo
// entries (one per declared HTTP method). The blueprint prefix is applied via
// bpPrefixes unless the Indico "!" app-root prefix was present; a registration
// with no methods defaults to a single GET route.
func addURLRuleToRouteInfo(reg addURLRuleReg, bpPrefixes blueprintPrefix) []routeInfo {
	prefix := ""
	if !reg.appRoot {
		prefix = lookupBlueprintPrefix(bpPrefixes, reg.file, reg.blueprintVar)
	}
	fullPath := combinePath(prefix, reg.rawPath)
	params := extractURLParams(fullPath)
	openAPIPath := flaskPathToOpenAPI(fullPath)

	methods := reg.methods
	if len(methods) == 0 {
		methods = []string{"GET"}
	}
	var routes []routeInfo
	for _, m := range methods {
		routes = append(routes, routeInfo{
			method:  m,
			path:    openAPIPath,
			handler: reg.handler,
			file:    reg.file,
			line:    reg.line,
			params:  params,
		})
	}
	return routes
}
