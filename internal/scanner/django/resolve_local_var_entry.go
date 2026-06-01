//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 단일 entry의 include(localVar) 참조를 로컬 인덱스 children으로 인라인 치환한다
package django

// resolveLocalVarEntry resolves a single entry's local-variable include. Existing
// inline children are resolved recursively first; then, if the entry references a
// known, non-cycling local list variable, the variable's children replace it as
// inline includes. The `visiting` set guards against self-referential variables.
func resolveLocalVarEntry(e urlEntry, index map[string][]urlEntry, visiting map[string]bool) urlEntry {
	if len(e.includeInline) > 0 {
		e.includeInline = resolveLocalVarIncludes(e.includeInline, index, visiting)
	}
	name := e.includeLocalVar
	children, ok := index[name]
	if !e.isInclude || name == "" || !ok || visiting[name] {
		return e
	}
	visiting[name] = true
	e.includeInline = resolveLocalVarIncludes(children, index, visiting)
	e.includeLocalVar = ""
	delete(visiting, name)
	return e
}
