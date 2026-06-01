//ff:func feature=scan type=parse control=selection topic=django
//ff:what 균형 괄호 스캔 1스텝: 현재 문자로 depth를 갱신하고(이스케이프는 건너뜀) 진행 인덱스를 반환한다
package django

// scanParenStep advances one character of a balanced-paren scan. It updates
// *depth for '(' / ')' and skips the char after a backslash escape, returning
// the (possibly advanced) index and the new depth.
func scanParenStep(runes []rune, i int, depth *int) (int, int) {
	switch runes[i] {
	case '\\':
		i++ // skip escaped char
	case '(':
		*depth++
	case ')':
		*depth--
	}
	return i, *depth
}
