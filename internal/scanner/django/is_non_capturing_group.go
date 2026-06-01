//ff:func feature=scan type=parse control=iteration dimension=1 topic=django
//ff:what 정규식 그룹 내부가 비포획 그룹(?:...) 또는 lookaround(?=,?!,?<=,?<!)인지 판정한다
package django

import "strings"

// isNonCapturingGroup reports whether a group's inner text starts with a
// non-capturing/lookaround marker "(?:...", "(?=...", "(?!...", "(?<=...",
// "(?<!...". Such groups never bind a path parameter.
func isNonCapturingGroup(inner string) bool {
	for _, p := range []string{"?:", "?=", "?!", "?<=", "?<!"} {
		if strings.HasPrefix(inner, p) {
			return true
		}
	}
	return false
}
