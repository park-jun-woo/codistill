//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 괄호 표현식(parenthesized_expression)의 내부 표현식을 찾아 평가한다(괄호 토큰 제외)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func evalParenthesized(node *sitter.Node, src []byte, imports map[string]string, referrerPath, projectRoot string) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "(" || child.Type() == ")" {
			continue
		}
		return evalPathExpr(child, src, imports, referrerPath, projectRoot)
	}
	return ""
}
