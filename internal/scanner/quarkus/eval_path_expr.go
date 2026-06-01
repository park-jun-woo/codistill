//ff:func feature=scan type=extract control=selection topic=quarkus
//ff:what @Path 인자 표현식(리터럴/상수참조/문자열결합)을 재귀 평가하여 경로 문자열을 만든다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func evalPathExpr(node *sitter.Node, src []byte, imports map[string]string, referrerPath, projectRoot string) string {
	if node == nil {
		return ""
	}
	switch node.Type() {
	case "string_literal":
		return unquoteJava(nodeText(node, src))
	case "identifier", "field_access", "scoped_identifier":
		return resolveConstantValue(nodeText(node, src), imports, referrerPath, projectRoot)
	case "binary_expression":
		return evalBinaryConcat(node, src, imports, referrerPath, projectRoot)
	case "parenthesized_expression":
		return evalParenthesized(node, src, imports, referrerPath, projectRoot)
	default:
		return ""
	}
}
