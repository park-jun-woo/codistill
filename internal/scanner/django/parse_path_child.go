//ff:func feature=scan type=extract control=selection topic=django
//ff:what 리스트의 직접 자식 노드 하나를 call/list_splat 타입에 따라 urlEntry로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parsePathChild dispatches a single direct child of a urlpatterns list to the
// appropriate parser based on its node type: `call` nodes go to parsePathCall
// (path()/re_path()) and `list_splat` nodes go to parseListSplat (`*router.urls`).
// Other node types yield nil.
func parsePathChild(child *sitter.Node, src []byte) *urlEntry {
	switch child.Type() {
	case "call":
		return parsePathCall(child, src)
	case "list_splat":
		return parseListSplat(child, src)
	}
	return nil
}
