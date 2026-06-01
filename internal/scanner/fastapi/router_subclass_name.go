//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what class_definition이 새로운 라우터 서브클래스면 그 이름을, 아니면 빈 문자열을 반환
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// routerSubclassName returns the class name of cls when it is a not-yet-known
// router subclass (directly or transitively inheriting from APIRouter), and an
// empty string otherwise.
func routerSubclassName(cls *sitter.Node, src []byte, subclasses map[string]bool) string {
	nameNode := findChildByType(cls, "identifier")
	if nameNode == nil {
		return ""
	}
	name := nodeText(nameNode, src)
	if name == "" || subclasses[name] {
		return ""
	}
	if !isRouterSubclassOf(collectParentNames(cls, src), subclasses) {
		return ""
	}
	return name
}
