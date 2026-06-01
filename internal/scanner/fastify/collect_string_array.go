//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 배열 노드에서 문자열 리터럴 요소만 문자열 슬라이스로 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectStringArray(arr *sitter.Node, src []byte) []string {
	var out []string
	for _, el := range collectArrayElements(arr) {
		if el.Type() == "string" {
			out = append(out, unquoteTS(nodeText(el, src)))
		}
	}
	return out
}
