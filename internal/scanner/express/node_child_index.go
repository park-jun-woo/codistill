//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 부모 노드에서 지정 자식 노드의 인덱스를 반환한다(없으면 -1)
package express

import sitter "github.com/smacker/go-tree-sitter"

// nodeChildIndex returns the index of child within parent, or -1 if not found.
func nodeChildIndex(parent, child *sitter.Node) int {
	for i := 0; i < int(parent.ChildCount()); i++ {
		if parent.Child(i) == child {
			return i
		}
	}
	return -1
}
