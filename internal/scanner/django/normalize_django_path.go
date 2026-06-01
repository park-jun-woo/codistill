//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what Django URL 정규식을 OpenAPI path 문자열과 path parameter 목록으로 완전 정규화한다
package django

// normalizeDjangoPath converts a Django URL pattern (path() or re_path() regex)
// into an OpenAPI-style path string and the ordered list of path parameters.
//
// It handles, at every position (not just prefix/suffix):
//   - named groups "(?P<name>...)"   -> "{name}"
//   - "<conv:name>" / "<name>"        -> "{name}"
//   - unnamed capture groups "(...)" -> "{paramN}" (N counts unnamed groups L->R)
//   - non-capturing groups "(?:...)" / lookarounds -> dropped
//   - regex anchors "^" and "$"       -> removed everywhere
//   - per-segment regex remnants ("\.", "\d", "+", "*", "[...]") cleaned
//     conservatively while pure-literal segments are preserved verbatim.
//
// The returned path and params are produced together so callers always agree on
// the synthesized "{paramN}" names for positional groups.
func normalizeDjangoPath(path string) (string, []urlParam) {
	n := &pathNormalizer{runes: []rune(path)}
	for i := 0; i < len(n.runes); i++ {
		i = n.step(i)
	}
	return cleanRegexRemnants(n.out.String()), n.params
}
