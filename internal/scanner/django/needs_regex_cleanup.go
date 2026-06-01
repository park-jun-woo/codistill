//ff:func feature=scan type=parse control=sequence topic=django
//ff:what 문자열에 정리가 필요한 정규식 메타문자(\ [ ] ( ) + * ? | { } 등)가 남아있는지 판정한다
package django

import "strings"

// needsRegexCleanup reports whether s still carries regex metacharacters that
// warrant cleanup. A lone "." is treated as a normal literal (filenames like
// "jwks.json") and does not trigger cleanup on its own.
func needsRegexCleanup(s string) bool {
	return strings.ContainsAny(s, `\[]()+*?|^$`)
}
