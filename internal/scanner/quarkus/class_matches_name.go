//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스 노드의 식별자가 주어진 이름과 일치하는지 판정한다(빈 이름이면 항상 일치)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func classMatchesName(cls *sitter.Node, src []byte, className string) bool {
	if className == "" {
		return true
	}
	nameNode := findChildByType(cls, "identifier")
	return nameNode != nil && nodeText(nameNode, src) == className
}
