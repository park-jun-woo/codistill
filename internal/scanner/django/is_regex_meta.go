//ff:func feature=scan type=parse control=sequence topic=django
//ff:what 룬이 드롭 대상 정규식 메타문자(+ * ? | [ ] ( ) ^ $)인지 판정한다
package django

import "strings"

// isRegexMeta reports whether c is a regex quantifier/grouping/class/anchor
// character that should be dropped during conservative segment cleanup.
func isRegexMeta(c rune) bool {
	return strings.ContainsRune(`+*?|[]()^$`, c)
}
