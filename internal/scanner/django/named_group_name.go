//ff:func feature=scan type=parse control=sequence topic=django
//ff:what 정규식 그룹 내부 문자열이 named group(?P<name>...)이면 이름을 추출한다
package django

import "strings"

// namedGroupName reports whether a group's inner text (the part between the
// outer parentheses) is a Python named group "?P<name>..." and returns name.
func namedGroupName(inner string) (string, bool) {
	const prefix = "?P<"
	if !strings.HasPrefix(inner, prefix) {
		return "", false
	}
	rest := inner[len(prefix):]
	end := strings.IndexByte(rest, '>')
	if end <= 0 {
		return "", false
	}
	name := rest[:end]
	if !isValidParamName(name) {
		return "", false
	}
	return name, true
}
