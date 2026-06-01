//ff:func feature=scan type=extract control=iteration dimension=2 topic=quarkus
//ff:what 클래스/인터페이스에서 static final(인터페이스 필드는 암묵적) 필드 값을 추출한다(리터럴/결합식 재귀 평가)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func findStaticFinalField(cls *sitter.Node, src []byte, fieldName string, imports map[string]string, referrerPath, projectRoot string) string {
	body := findChildByType(cls, "class_body")
	implicitStaticFinal := false
	if body == nil {
		body = findChildByType(cls, "interface_body")
		implicitStaticFinal = true
	}
	if body == nil {
		return ""
	}
	for i := 0; i < int(body.ChildCount()); i++ {
		field := body.Child(i)
		if field.Type() != "field_declaration" && field.Type() != "constant_declaration" {
			continue
		}
		if !implicitStaticFinal && !hasModifiers(field, src, "static", "final") {
			continue
		}
		for j := 0; j < int(field.ChildCount()); j++ {
			decl := field.Child(j)
			if decl.Type() != "variable_declarator" {
				continue
			}
			nameNode := findChildByType(decl, "identifier")
			if nameNode == nil || nodeText(nameNode, src) != fieldName {
				continue
			}
			return evalDeclaratorValue(decl, src, imports, referrerPath, projectRoot)
		}
	}
	return ""
}
