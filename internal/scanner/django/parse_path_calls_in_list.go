//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 노드의 직접 자식 call에서 path()/re_path() 호출을 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parsePathCallsInList parses path()/re_path() calls among a node's direct children
// (used for both list literals and wrapper argument lists like i18n_patterns(...)).
// `list_splat` children (e.g. `*router.urls`) are dispatched to parseListSplat so a
// router's registered ViewSets are wired into the urlconf.
func parsePathCallsInList(listNode *sitter.Node, src []byte) []urlEntry {
	var entries []urlEntry
	for i := 0; i < int(listNode.ChildCount()); i++ {
		if entry := parsePathChild(listNode.Child(i), src); entry != nil {
			entries = append(entries, *entry)
		}
	}
	return entries
}
