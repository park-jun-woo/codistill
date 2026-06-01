//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what Flask-AppBuilder @expose/ModelRestApi 라우트를 전 파일에서 추출한다
package flask

// extractAppbuilderRoutes resolves Flask-AppBuilder API routes. It collects every
// BaseApi / ModelRestApi / *RestApi subclass (collectAppbuilderAPIs) and, per
// class, emits the explicit @expose("/path", methods=...) routes
// (exposeRoutesFromClass) plus — for ModelRestApi-family classes — the standard
// CRUD endpoints (expandModelRestApi), each composed against the class base_url.
// Explicit @expose paths suppress the synthesized CRUD entry on the same path so
// overrides win. Registration via appbuilder.add_api is not required: a class
// definition alone yields routes (recall-first per Phase172).
func extractAppbuilderRoutes(files []fileInfo) []routeInfo {
	apis := collectAppbuilderAPIs(files)
	if len(apis) == 0 {
		return nil
	}
	aliasesByFile := make(map[string]importAlias, len(files))
	for _, fi := range files {
		aliasesByFile[fi.relPath] = collectImportAliases(fi.root, fi.src)
	}
	srcByFile := make(map[string][]byte, len(files))
	for _, fi := range files {
		srcByFile[fi.relPath] = fi.src
	}

	var routes []routeInfo
	for _, api := range apis {
		src := srcByFile[api.file]
		aliases := aliasesByFile[api.file]
		exposeRoutes, exposed := exposeRoutesFromClass(api, aliases, src)
		routes = append(routes, exposeRoutes...)
		if api.isModelRestApi {
			routes = append(routes, expandModelRestApi(api, exposed)...)
		}
	}
	return routes
}
