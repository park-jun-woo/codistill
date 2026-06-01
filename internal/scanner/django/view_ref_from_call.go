//ff:func feature=scan type=extract control=sequence topic=django
//ff:what call 노드를 view 참조로 해석하고 .as_view 접미사를 제거한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// viewRefFromCall resolves a call node to a view reference, stripping a
// trailing .as_view from attribute callees.
func viewRefFromCall(node *sitter.Node, src []byte) string {
	if attr := findChildByType(node, "attribute"); attr != nil {
		return strings.TrimSuffix(nodeText(attr, src), ".as_view")
	}
	if id := findChildByType(node, "identifier"); id != nil {
		return nodeText(id, src)
	}
	return ""
}
