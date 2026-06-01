//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 자식 노드가 import된 이름인지 확인하여 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryImportedName returns the imported name if this child is an imported identifier.
func tryImportedName(child *sitter.Node, stmt *sitter.Node, idx int, src []byte) string {
	if idx == 0 {
		return ""
	}
	prev := stmt.Child(idx - 1)
	prevText := nodeText(prev, src)

	// 괄호 묶음 import(`from . import (a, b)`)에서 첫 이름은 직전 토큰이 "("다.
	if prevText != "import" && prevText != "," && prevText != "(" {
		return ""
	}
	switch child.Type() {
	case "dotted_name", "identifier":
		return nodeText(child, src)
	}
	return ""
}
