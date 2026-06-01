//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what add_url_rule 등록을 prefix와 결합해 라우트(routeInfo)로 추출한다
package flask

// extractAddURLRuleRoutes resolves routes registered via
// X.add_url_rule(rule, endpoint, view, methods=...) across all files. This is
// the registration form Indico's IndicoBlueprint uses to attach RequestHandler
// classes (and it also covers standard Flask add_url_rule with a function
// view). Per-registration prefix composition and method expansion is delegated
// to addURLRuleToRouteInfo.
func extractAddURLRuleRoutes(files []fileInfo, bpPrefixes blueprintPrefix) []routeInfo {
	var routes []routeInfo
	for _, fi := range files {
		for _, reg := range collectAddURLRule(fi.root, fi.src, fi.relPath) {
			routes = append(routes, addURLRuleToRouteInfo(reg, bpPrefixes)...)
		}
	}
	return routes
}
