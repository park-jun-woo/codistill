//ff:func feature=scan type=convert control=sequence topic=django
//ff:what 백슬래시 이스케이프 처리: 다음 문자를 리터럴로 출력하고(백슬래시 제거) 소비 인덱스를 반환한다
package django

import "strings"

// writeEscaped handles a backslash at runes[i]. If a char follows, it is written
// as a literal (the backslash dropped) and i+1 is returned; otherwise a trailing
// backslash is dropped and i is returned.
func writeEscaped(b *strings.Builder, runes []rune, i int) int {
	if i+1 >= len(runes) {
		return i // dangling backslash: drop
	}
	b.WriteRune(runes[i+1])
	return i + 1
}
