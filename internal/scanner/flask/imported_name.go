//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what import된 이름 노드에서 (로컬명, 원본명)을 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// importedName extracts (localName, originalName) from one imported-name node.
// A plain dotted_name binds the name as-is; an aliased_import binds the alias.
func importedName(n *sitter.Node, src []byte) (string, string) {
	if n.Type() == "aliased_import" {
		nameNode := n.ChildByFieldName("name")
		aliasNode := n.ChildByFieldName("alias")
		if nameNode == nil || aliasNode == nil {
			return "", ""
		}
		return nodeText(aliasNode, src), nodeText(nameNode, src)
	}
	if n.Type() == "dotted_name" {
		t := nodeText(n, src)
		return t, t
	}
	return "", ""
}
