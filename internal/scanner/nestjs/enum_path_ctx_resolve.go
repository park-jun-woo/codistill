//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what enumPathCtx로 단일 경로 인자의 enum 멤버표현식을 해석한다
package nestjs

// resolve returns the enum-resolved value of arg (Enum.Member), falling back to
// const string-identifier resolution (simple CONST reference), then to
// template-literal `${ident}` interpolation, or arg unchanged when none apply.
// Member expressions go through resolveEnumPathArg; simple identifiers through
// resolveConstStringIdentifier — the two branches are mutually exclusive (dot
// vs. no dot), so no double resolution occurs. interpolateTemplatePath only
// fires when arg contains a `${...}` fragment (e.g.
// `${PREFIX}/:modelId/records`), substituting each resolvable fragment.
func (pc enumPathCtx) resolve(arg string) string {
	if resolved, ok := resolveEnumPathArg(arg, pc.root, pc.src, pc.absFile, pc.imports, pc.projectRoot); ok {
		return resolved
	}
	if resolved, ok := resolveConstStringIdentifier(arg, pc.root, pc.src, pc.absFile, pc.imports, pc.projectRoot); ok {
		return resolved
	}
	if resolved, ok := pc.interpolateTemplatePath(arg); ok {
		return resolved
	}
	return arg
}
