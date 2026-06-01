//ff:func feature=scan type=parse control=iteration dimension=1 topic=django
//ff:what 문자열이 유효한 파라미터/컨버터 식별자([A-Za-z_][A-Za-z0-9_]*)인지 판정한다
package django

// isValidParamName reports whether s is a valid Python identifier usable as a
// Django path variable or converter name: [A-Za-z_][A-Za-z0-9_]*.
func isValidParamName(s string) bool {
	if s == "" || !isIdentStart(s[0]) {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !isIdentPart(s[i]) {
			return false
		}
	}
	return true
}
