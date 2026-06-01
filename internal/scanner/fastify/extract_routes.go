//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 파일에서 fastify.get/post/put/patch/delete 라우트 호출을 추출한다
package fastify

func extractRoutes(fi *fileInfo, instances map[string]bool) []routeInfo {
	var routes []routeInfo
	for _, call := range findAllByType(fi.Root, "call_expression") {
		ri := extractOneRoute(call, fi.Src, instances)
		if ri != nil {
			routes = append(routes, *ri)
			continue
		}
		routes = append(routes, extractObjectRoute(call, fi.Src, instances)...)
	}
	return routes
}
