//ff:func feature=scan type=extract control=selection topic=express
//ff:what arguments 첫 인자가 string이면 1건, array면 string 요소마다 동일 handler로 routeInfo를 다건 생성한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func buildRoutesFromArgs(args *sitter.Node, src []byte, method string, line int) []routeInfo {
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return nil
	}
	pathNode := argNodes[0]
	switch pathNode.Type() {
	case "string":
		path := unquoteTS(nodeText(pathNode, src))
		ri := buildRouteWithPath(argNodes, src, method, path, line)
		return []routeInfo{*ri}
	case "array":
		return buildRoutesFromArrayPath(pathNode, argNodes, src, method, line)
	default:
		return nil
	}
}
