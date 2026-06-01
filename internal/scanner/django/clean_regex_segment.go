//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what 단일 path 세그먼트에서 정규식 잔재를 제거해 리터럴 텍스트로 보수 정리한다
package django

import "strings"

// cleanRegexSegment conservatively strips regex syntax from a single path
// segment that is known to contain metacharacters.
//
//   - "\X"  -> "X"   (escaped literal, e.g. "\." -> ".", "\-" -> "-")
//   - regex quantifier / grouping / class chars ("+ * ? | [ ] ( ) ^ $") dropped
//
// Remaining characters (letters, digits, "." "-" "_" etc.) are preserved so a
// segment like "\.well-known" becomes ".well-known".
func cleanRegexSegment(seg string) string {
	var b strings.Builder
	runes := []rune(seg)
	for i := 0; i < len(runes); i++ {
		i = cleanRegexChar(&b, runes, i)
	}
	return b.String()
}
