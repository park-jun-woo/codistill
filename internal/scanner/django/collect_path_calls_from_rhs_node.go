//ff:func feature=scan type=extract control=selection topic=django
//ff:what 괄호/연결 RHS 노드 타입을 분기해 list면 path() 수집, 래퍼면 자식 재귀로 위임한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// collectPathCallsFromRHSNode walks a urlpatterns RHS expression that may be wrapped in
// parentheses or composed with `+` (e.g. `( [path(...), ...] + extra + static(...) )`).
// Only `list` operands contribute path()/re_path() entries; variable/function-call operands
// (e.g. `plugin_registry.urls`, `static(...)`) are ignored to stay regression-safe.
func collectPathCallsFromRHSNode(node *sitter.Node, src []byte) []urlEntry {
	if node == nil {
		return nil
	}
	switch node.Type() {
	case "list":
		return parsePathCallsInList(node, src)
	case "parenthesized_expression", "binary_operator":
		return collectPathCallsFromRHSChildren(node, src)
	}
	return nil
}
