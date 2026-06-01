//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what 템플릿 조각 표현식을 Enum.Member 또는 단순 const로 해석한다
package nestjs

// resolveTemplateFragment resolves a single `${...}` template fragment
// expression as an Enum.Member (resolveEnumPathArg) or a simple const
// identifier (resolveConstStringIdentifier), mirroring pc.resolve's branches.
// Returns ("", false) for anything not statically resolvable (function calls,
// arithmetic, etc.), so the caller keeps the fragment verbatim.
func (pc enumPathCtx) resolveTemplateFragment(expr string) (string, bool) {
	if resolved, ok := resolveEnumPathArg(expr, pc.root, pc.src, pc.absFile, pc.imports, pc.projectRoot); ok {
		return resolved, true
	}
	if resolved, ok := resolveConstStringIdentifier(expr, pc.root, pc.src, pc.absFile, pc.imports, pc.projectRoot); ok {
		return resolved, true
	}
	return "", false
}
