//ff:func feature=scan type=extract control=selection topic=django
//ff:what 단일 노드를 view 참조 문자열(identifier/attribute/.as_view 호출)로 해석한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// viewRefFromNode resolves a single argument-value node to a view reference.
// Identifiers and attributes map to their text; a .as_view() call maps to the
// class name. Other node types yield "".
func viewRefFromNode(node *sitter.Node, src []byte) string {
	switch node.Type() {
	case "identifier", "attribute":
		return nodeText(node, src)
	case "call":
		return viewRefFromCall(node, src)
	default:
		return ""
	}
}
