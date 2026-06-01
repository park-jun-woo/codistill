//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 lexical_declaration에서 이름이 name인 string 리터럴 value를 해소한다 (없으면 "")
package express

import sitter "github.com/smacker/go-tree-sitter"

func constStringInDecl(decl *sitter.Node, src []byte, name string) string {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		nameNode := declarator.ChildByFieldName("name")
		if nameNode == nil || nodeText(nameNode, src) != name {
			continue
		}
		value := declarator.ChildByFieldName("value")
		if value != nil && value.Type() == "string" {
			return unquoteTS(nodeText(value, src))
		}
	}
	return ""
}
