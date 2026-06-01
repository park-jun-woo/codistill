//ff:func feature=scan type=convert control=selection topic=django
//ff:what 세그먼트 정리 1스텝: 이스케이프 리터럴은 백슬래시 제거 후 출력, 메타문자는 드롭, 나머지는 출력
package django

import "strings"

// cleanRegexChar processes runes[i] for cleanRegexSegment, writing the cleaned
// output to b and returning the index actually consumed (advanced past an
// escaped char when runes[i] is a backslash).
func cleanRegexChar(b *strings.Builder, runes []rune, i int) int {
	c := runes[i]
	switch {
	case c == '\\':
		return writeEscaped(b, runes, i)
	case isRegexMeta(c):
		return i // drop metacharacter
	default:
		b.WriteRune(c)
		return i
	}
}
