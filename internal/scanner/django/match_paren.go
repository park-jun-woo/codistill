//ff:func feature=scan type=parse control=iteration dimension=1 topic=django
//ff:what 정규식 문자열에서 여는 괄호에 대응하는 닫는 괄호 인덱스를 균형 매칭으로 찾는다
package django

// matchParen returns the index of the ')' that closes the '(' at runes[open],
// honoring nesting and "\(" escapes. Returns -1 if unbalanced.
func matchParen(runes []rune, open int) int {
	depth := 0
	for i := open; i < len(runes); i++ {
		i, depth = scanParenStep(runes, i, &depth)
		if depth == 0 {
			return i
		}
	}
	return -1
}
