//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what array path 노드의 string 요소마다 동일 handler/middleware로 routeInfo를 1건씩 생성한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func buildRoutesFromArrayPath(arr *sitter.Node, argNodes []*sitter.Node, src []byte, method string, line int) []routeInfo {
	var routes []routeInfo
	for _, path := range collectArrayStringPaths(arr, src) {
		ri := buildRouteWithPath(argNodes, src, method, path, line)
		routes = append(routes, *ri)
	}
	return routes
}
