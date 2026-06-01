//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what keyword_argument 노드의 값 노드를 view 참조 문자열로 해석한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// keywordArgViewValue returns the view reference of a keyword_argument node's
// value (the right-hand side of `name=value`), given the already-located key
// identifier node. The value may be an identifier (get_messages), an attribute
// (views.get_messages) or a .as_view() call (resolved to the class name).
func keywordArgViewValue(kwNode, keyNode *sitter.Node, src []byte) string {
	for i := 0; i < int(kwNode.ChildCount()); i++ {
		child := kwNode.Child(i)
		if child == keyNode {
			continue
		}
		if view := viewRefFromNode(child, src); view != "" {
			return view
		}
	}
	return ""
}
