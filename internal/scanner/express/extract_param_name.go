//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 경로 세그먼트에서 :param 이름을 추출한다
package express

import "strings"

func extractParamName(part string) string {
	if !strings.HasPrefix(part, ":") {
		return ""
	}
	name := part[1:]
	if idx := strings.Index(name, "("); idx >= 0 {
		name = name[:idx]
	}
	name = strings.TrimSuffix(name, "?")
	return name
}
