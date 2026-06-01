//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 메서드 목록과 공통 라우트 속성으로 객체형 routeInfo 목록을 생성한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func buildObjectRoutes(methods []string, path, handler string, line int, startByte uint32, opts *sitter.Node) []routeInfo {
	var routes []routeInfo
	for _, method := range methods {
		routes = append(routes, routeInfo{
			Method:    method,
			Path:      path,
			Handler:   handler,
			Line:      line,
			StartByte: startByte,
			Schema:    opts,
		})
	}
	return routes
}
