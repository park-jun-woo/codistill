//ff:func feature=scan type=extract control=sequence topic=express
//ff:what variable_declarator의 init이 require('...') 호출이면 모듈 경로 문자열을 반환한다 (아니면 "")
package express

import sitter "github.com/smacker/go-tree-sitter"

func requirePathOfDeclarator(declarator *sitter.Node, src []byte) string {
	callNode := findInitValue(declarator)
	if callNode == nil || callNode.Type() != "call_expression" {
		return ""
	}
	fnNode := findChildByType(callNode, "identifier")
	if fnNode == nil || nodeText(fnNode, src) != "require" {
		return ""
	}
	return extractRequirePath(callNode, src)
}
