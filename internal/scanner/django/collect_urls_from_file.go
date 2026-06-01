//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 단일 파일에서 urlpatterns를 수집한다
package django

// collectURLsFromFile extracts urlpatterns from a single file. include(localVar)
// references to file-scoped local list variables are resolved inline against the
// file's local list-variable index before returning.
func collectURLsFromFile(fi fileInfo) []urlEntry {
	var entries []urlEntry
	entries = append(entries, collectFromAssignments(fi)...)
	entries = append(entries, collectFromAugmentedAssignments(fi)...)
	entries = append(entries, collectStarImportIncludes(fi)...)
	return resolveLocalVarIncludes(entries, collectLocalListVars(fi), map[string]bool{})
}
