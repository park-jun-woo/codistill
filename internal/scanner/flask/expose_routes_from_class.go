//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 한 Flask-AppBuilder API 클래스에서 @expose 메서드 라우트를 추출한다
package flask

// exposeRoutesFromClass extracts the explicit @expose("/path", methods=...)
// routes declared on the methods of one Flask-AppBuilder API class. Each
// decorated method is delegated to exposeRoutesFromMethod, which composes the
// path against api.baseURL. The returned map of OpenAPI paths lets callers let
// explicit @expose overrides win over synthesized ModelRestApi CRUD.
func exposeRoutesFromClass(api appbuilderAPIInfo, aliases importAlias, src []byte) ([]routeInfo, map[string]bool) {
	body := findChildByType(api.node, "block")
	if body == nil {
		return nil, nil
	}
	var routes []routeInfo
	exposed := make(map[string]bool)
	for _, dd := range childrenOfType(body, "decorated_definition") {
		mr := exposeRoutesFromMethod(dd, api, aliases, src)
		routes = append(routes, mr...)
		for _, r := range mr {
			exposed[r.path] = true
		}
	}
	return routes, exposed
}
