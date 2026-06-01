//ff:func feature=scan type=parse control=sequence topic=django
//ff:what Django <conv:name> 변수 표기에서 컨버터 타입(int/str/uuid/slug/path)을 추출한다
package django

import "strings"

// angleConverter extracts the converter type (e.g. "int", "slug") from the
// inside of a Django "<conv:name>" token. Returns "" when no converter prefix
// is present (plain "<name>").
func angleConverter(inner string) string {
	i := strings.IndexByte(inner, ':')
	if i < 0 {
		return ""
	}
	conv := inner[:i]
	if !isValidParamName(conv) {
		return ""
	}
	return conv
}
