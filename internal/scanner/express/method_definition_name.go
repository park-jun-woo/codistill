//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 메서드 정의 노드에서 메서드명(핸들러명)을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// methodDefinitionName returns the handler name of a method_definition node,
// or "" when no property_identifier is found.
func methodDefinitionName(m *sitter.Node, src []byte) string {
	name := findChildByType(m, "property_identifier")
	if name == nil {
		return ""
	}
	return nodeText(name, src)
}
