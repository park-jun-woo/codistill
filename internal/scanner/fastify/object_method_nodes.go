//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what 객체형 라우트의 method pair 값을 문자열 1개 또는 배열 요소 목록으로 분해한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func objectMethodNodes(methodNode *sitter.Node) []*sitter.Node {
	switch methodNode.Type() {
	case "string", "template_string":
		return []*sitter.Node{methodNode}
	case "array":
		return collectArrayElements(methodNode)
	default:
		return nil
	}
}
