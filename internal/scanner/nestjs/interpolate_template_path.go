//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 템플릿리터럴 경로의 ${ident} 조각을 enum/const로 해석해 보간 치환한다
package nestjs

import "strings"

// interpolateTemplatePath resolves `${ident}` / `${Enum.Member}` fragments
// embedded in a template-literal decorator path (e.g.
// `${PREFIX_APIV3_DATA}/:modelId/records`) using the same enum/const resolution
// assets as pc.resolve. Each `${...}` fragment is scanned without regex; the
// inner expression is resolved via resolveEnumPathArg (Enum.Member) or
// resolveConstStringIdentifier (simple const). Fragments that cannot be
// statically resolved (function calls, arithmetic, etc.) are left verbatim so
// the 2nd-line guard can strip fake params downstream. Returns (result, true)
// only when the path actually contained a `${...}` fragment.
func (pc enumPathCtx) interpolateTemplatePath(path string) (string, bool) {
	if !strings.Contains(path, "${") {
		return path, false
	}
	var b strings.Builder
	rest := path
	found := false
	for {
		open := strings.Index(rest, "${")
		if open < 0 {
			b.WriteString(rest)
			break
		}
		close := strings.IndexByte(rest[open:], '}')
		if close < 0 {
			b.WriteString(rest)
			break
		}
		found = true
		b.WriteString(rest[:open])
		expr := strings.TrimSpace(rest[open+2 : open+close])
		if resolved, ok := pc.resolveTemplateFragment(expr); ok {
			b.WriteString(resolved)
		} else {
			// keep the verbatim ${...} fragment for the downstream guard
			b.WriteString(rest[open : open+close+1])
		}
		rest = rest[open+close+1:]
	}
	if !found {
		return path, false
	}
	return b.String(), true
}
