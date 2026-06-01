//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urlpatterns += <rhs> 증분 대입문(list/래퍼 call/괄호)에서 path() 호출을 수집한다
package django

// collectFromAugmentedAssignments extracts urlEntries from urlpatterns += <rhs>
// assignments. The RHS may be a list literal (`urlpatterns += [path(...)]`), a
// wrapper call (`urlpatterns += i18n_patterns(path(...), ...)`), or a
// parenthesized/concatenated expression; all are dispatched through
// collectFromURLPatternsRHS so trailing routes added via augmented assignment
// (including those inside `if DEBUG:` branches, since findAllByType recurses the
// whole tree) are not dropped.
func collectFromAugmentedAssignments(fi fileInfo) []urlEntry {
	var entries []urlEntry
	for _, node := range findAllByType(fi.root, "augmented_assignment") {
		leftNode := findChildByType(node, "identifier")
		if leftNode == nil || nodeText(leftNode, fi.src) != "urlpatterns" {
			continue
		}
		entries = append(entries, collectFromURLPatternsRHS(node, fi.src)...)
	}
	return entries
}
