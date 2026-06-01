//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what string_literal 노드에서 string_content 내용을 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func stringLiteralContent(strLit *sitter.Node, src []byte) string {
	strContent := findChildByType(strLit, "string_content")
	if strContent == nil {
		return ""
	}
	return nodeText(strContent, src)
}
