//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 단일 call_expression에서 Fastify 라우트 정보를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractOneRoute(call *sitter.Node, src []byte, instances map[string]bool) *routeInfo {
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return nil
	}
	obj := routeReceiverIdent(fn)
	if obj == nil || !instances[nodeText(obj, src)] {
		return nil
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	method, ok := httpMethods[nodeText(prop, src)]
	if !ok {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 2 {
		return nil
	}
	pathNode := argNodes[0]
	if pathNode.Type() != "string" && pathNode.Type() != "template_string" {
		return nil
	}
	ri := &routeInfo{
		Method:    method,
		Path:      unquoteTS(nodeText(pathNode, src)),
		Line:      int(call.StartPoint().Row) + 1,
		StartByte: call.StartByte(),
	}
	assignSchemaAndHandler(ri, argNodes, src)
	return ri
}
