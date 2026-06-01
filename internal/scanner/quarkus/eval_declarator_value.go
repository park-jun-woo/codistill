//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what variable_declarator의 초기화 표현식을 평가한다(리터럴/결합식/상수참조)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func evalDeclaratorValue(decl *sitter.Node, src []byte, imports map[string]string, referrerPath, projectRoot string) string {
	for j := 0; j < int(decl.ChildCount()); j++ {
		val := decl.Child(j)
		switch val.Type() {
		case "identifier", "=":
			continue
		case "decimal_integer_literal", "decimal_floating_point_literal", "true", "false":
			return nodeText(val, src)
		default:
			return evalPathExpr(val, src, imports, referrerPath, projectRoot)
		}
	}
	return ""
}
