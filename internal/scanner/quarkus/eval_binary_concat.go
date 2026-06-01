//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 문자열 결합(binary_expression)의 각 피연산자를 평가하여 이어 붙인다('+' 토큰 제외)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func evalBinaryConcat(node *sitter.Node, src []byte, imports map[string]string, referrerPath, projectRoot string) string {
	var out string
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "+" {
			continue
		}
		out += evalPathExpr(child, src, imports, referrerPath, projectRoot)
	}
	return out
}
