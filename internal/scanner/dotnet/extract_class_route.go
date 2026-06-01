//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 클래스의 [Route] 어트리뷰트에서 prefix 경로를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractClassRoute(cls *sitter.Node, src []byte, className string) string {
	attr := findAttribute(cls, src, AttrRoute)
	if attr == nil {
		return ""
	}
	route := attributeFirstStringArg(attr, src)
	if route == "" {
		return ""
	}
	expanded := expandRouteTokens(route, className, "")
	// 클래스-레벨 [Route("~/...")] / [Route("/...")] 는 앱 루트 기준 절대 prefix이므로
	// 선두 절대 마커(~/·/)를 제거한다. joinPath 는 / 만 trim 하므로 ~ 가 리터럴
	// 세그먼트로 남는 것을 방지한다. (상대 prefix 는 마커가 없어 그대로 통과)
	stripped, _ := isAbsoluteRoute(expanded)
	return stripped
}
