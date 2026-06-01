//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what include(localVar) 엔트리를 파일 로컬 리스트변수 인덱스의 entry 묶음으로 인라인 치환한다
package django

// resolveLocalVarIncludes rewrites entries that reference a file-scoped local list
// variable via include(varName) into inline includes carrying the variable's parsed
// children. The same variable referenced under several prefixes (e.g. api/v1/, api/v2/,
// api/v3/) is expanded independently at each call site because each entry receives its
// own copy of the children. The `visiting` set guards against self-referential local
// variables. Entries with an unresolved local-var name are left as-is (dropped later).
func resolveLocalVarIncludes(entries []urlEntry, index map[string][]urlEntry, visiting map[string]bool) []urlEntry {
	out := make([]urlEntry, 0, len(entries))
	for _, e := range entries {
		out = append(out, resolveLocalVarEntry(e, index, visiting))
	}
	return out
}
