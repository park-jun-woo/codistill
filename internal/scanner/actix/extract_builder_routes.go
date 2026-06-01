//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 빌더 패턴(web::resource().route(), web::scope().service())에서 라우트를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractBuilderRoutes(fi *fileInfo, handlerFuncs map[string]*handlerInfo) []builderRoute {
	var routes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectTopLevelServiceCall(n, fi, &routes, handlerFuncs)
		collectTopLevelRouteCall(n, fi, &routes)
	})
	return deduplicateBuilderRoutes(routes)
}
