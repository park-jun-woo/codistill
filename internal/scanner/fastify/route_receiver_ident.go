//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what member_expression의 object를 체인 끝까지 따라가 루트 인스턴스 식별자 노드를 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func routeReceiverIdent(fn *sitter.Node) *sitter.Node {
	obj := fn.ChildByFieldName("object")
	for obj != nil {
		switch obj.Type() {
		case "identifier":
			return obj
		case "call_expression":
			obj = obj.ChildByFieldName("function")
		case "member_expression":
			obj = obj.ChildByFieldName("object")
		default:
			return nil
		}
	}
	return nil
}
