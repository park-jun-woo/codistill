//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what ModelRestApi 서브클래스의 표준 CRUD 5엔드포인트를 합성한다
package flask

// expandModelRestApi synthesizes Flask-AppBuilder's standard ModelRestApi CRUD
// endpoints for one API class: GET <base>/ (list), GET <base>/<pk> (show),
// POST <base>/ (add), PUT <base>/<pk> (edit), DELETE <base>/<pk> (delete).
// Paths are composed against api.baseURL and normalized to OpenAPI form. An
// endpoint is skipped when an explicit @expose route (in exposed, keyed by
// OpenAPI path) already covers that path, so overrides win. The handler label
// is "ClassName.<op>".
func expandModelRestApi(api appbuilderAPIInfo, exposed map[string]bool) []routeInfo {
	type crud struct {
		method  string
		rawPath string
		op      string
	}
	specs := []crud{
		{"GET", "/", "get_list"},
		{"GET", "/<pk>", "get"},
		{"POST", "/", "post"},
		{"PUT", "/<pk>", "put"},
		{"DELETE", "/<pk>", "delete"},
	}
	var routes []routeInfo
	for _, s := range specs {
		fullPath := combinePath(api.baseURL, s.rawPath)
		openAPIPath := flaskPathToOpenAPI(fullPath)
		if exposed[openAPIPath] {
			continue
		}
		routes = append(routes, routeInfo{
			method:  s.method,
			path:    openAPIPath,
			handler: api.name + "." + s.op,
			file:    api.file,
			line:    int(api.node.StartPoint().Row) + 1,
			params:  extractURLParams(fullPath),
		})
	}
	return routes
}
