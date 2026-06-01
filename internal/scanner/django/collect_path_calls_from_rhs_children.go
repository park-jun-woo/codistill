//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 괄호/연결 노드의 모든 자식을 순회하며 RHS 재귀 수집기를 적용해 path() 엔트리를 모은다
package django

import sitter "github.com/smacker/go-tree-sitter"

// collectPathCallsFromRHSChildren recurses into every child of a parenthesized_expression
// or binary_operator node, aggregating path()/re_path() entries from any nested list operands.
func collectPathCallsFromRHSChildren(node *sitter.Node, src []byte) []urlEntry {
	var entries []urlEntry
	for i := 0; i < int(node.ChildCount()); i++ {
		entries = append(entries, collectPathCallsFromRHSNode(node.Child(i), src)...)
	}
	return entries
}
