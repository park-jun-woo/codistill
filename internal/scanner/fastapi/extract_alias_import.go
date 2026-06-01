//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what aliased_import 노드에서 (alias명, 원본명) 쌍을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractAliasImport extracts the alias name and original name from an
// aliased_import node (`import X as Y`). The original name is child 0
// (dotted_name/identifier) and the alias is the trailing identifier after `as`.
// Returns ("", "") if the node is not a well-formed aliased_import.
func extractAliasImport(node *sitter.Node, src []byte) (alias, origName string) {
	if node.Type() != "aliased_import" {
		return "", ""
	}
	orig := node.Child(0)
	if orig == nil {
		return "", ""
	}
	origName = nodeText(orig, src)
	for i := int(node.ChildCount()) - 1; i > 0; i-- {
		c := node.Child(i)
		if c.Type() == "identifier" {
			alias = nodeText(c, src)
			break
		}
	}
	if alias == "" || origName == "" {
		return "", ""
	}
	return alias, origName
}
