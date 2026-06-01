//ff:func feature=scan type=extract control=selection topic=dotnet
//ff:what 메서드 라우트 템플릿의 절대성(~/ · /)을 판정하고 선두 마커를 제거한 경로를 반환한다
package dotnet

import "strings"

func isAbsoluteRoute(template string) (string, bool) {
	switch {
	case strings.HasPrefix(template, "~/"):
		return template[2:], true
	case template == "~":
		return "", true
	case strings.HasPrefix(template, "/"):
		return template[1:], true
	default:
		return template, false
	}
}
