//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what 객체형 라우트의 handler pair 값 노드에서 핸들러 이름(identifier 또는 (anonymous))을 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func objectHandlerName(handlerNode *sitter.Node, src []byte) string {
	switch handlerNode.Type() {
	case "identifier":
		return nodeText(handlerNode, src)
	case "arrow_function", "function":
		return "(anonymous)"
	default:
		return ""
	}
}
