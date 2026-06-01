//ff:func feature=scan type=parse control=sequence topic=django
//ff:what Django <conv:name> 또는 <name> 변수 표기의 내부에서 파라미터 이름을 추출한다
package django

import "strings"

// angleParamName extracts the variable name from the inside of a Django
// "<conv:name>" or "<name>" token (the text between '<' and '>').
// Returns ok=false when inner is not a valid Django path variable.
func angleParamName(inner string) (string, bool) {
	name := inner
	if i := strings.IndexByte(inner, ':'); i >= 0 {
		name = inner[i+1:]
	}
	if !isValidParamName(name) {
		return "", false
	}
	return name, true
}
