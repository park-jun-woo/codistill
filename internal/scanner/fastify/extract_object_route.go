//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what server.route({ method, url, schema }) 객체형 등록에서 라우트 정보 목록을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractObjectRoute(call *sitter.Node, src []byte, instances map[string]bool) []routeInfo {
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return nil
	}
	obj := routeReceiverIdent(fn)
	if obj == nil || !instances[nodeText(obj, src)] {
		return nil
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil || nodeText(prop, src) != "route" {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 || argNodes[0].Type() != "object" {
		return nil
	}
	opts := argNodes[0]
	pathNode := findPairValue(opts, src, "url")
	if pathNode == nil || (pathNode.Type() != "string" && pathNode.Type() != "template_string") {
		return nil
	}
	methodNode := findPairValue(opts, src, "method")
	if methodNode == nil {
		return nil
	}
	methods := objectRouteMethods(methodNode, src)
	if len(methods) == 0 {
		return nil
	}
	handler := ""
	if h := findPairValue(opts, src, "handler"); h != nil {
		handler = objectHandlerName(h, src)
	}
	return buildObjectRoutes(methods, unquoteTS(nodeText(pathNode, src)), handler, int(call.StartPoint().Row)+1, call.StartByte(), opts)
}
