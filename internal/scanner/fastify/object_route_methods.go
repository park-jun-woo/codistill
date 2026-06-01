//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 객체형 라우트의 method pair 값(문자열 또는 배열)에서 정규화된 HTTP 메서드 목록을 추출한다
package fastify

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func objectRouteMethods(methodNode *sitter.Node, src []byte) []string {
	var methods []string
	for _, n := range objectMethodNodes(methodNode) {
		if n.Type() != "string" && n.Type() != "template_string" {
			continue
		}
		key := strings.ToLower(unquoteTS(nodeText(n, src)))
		if m, ok := httpMethods[key]; ok {
			methods = append(methods, m)
		}
	}
	return methods
}
