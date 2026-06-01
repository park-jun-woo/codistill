//ff:func feature=scan type=extract control=selection topic=quarkus
//ff:what 파라미터 어노테이션에 따라 path/query/header/form/body로 분류한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func classifyParam(param *sitter.Node, src []byte, ep *endpointInfo, imports map[string]string, referrerPath, projectRoot string) {
	typeName := extractParamType(param, src)
	paramName := extractParamName(param, src)

	switch {
	case hasAnnotation(param, src, AnnPathParam):
		classifyPathParam(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnQueryParam):
		classifyQueryParam(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnHeaderParam):
		classifyHeaderParam(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnFormParam):
		classifyFormParam(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnRestForm):
		classifyRestForm(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnBeanParam):
		ep.formType = typeName
	case hasAnnotation(param, src, AnnContext):
		// @Context 주입 객체(HttpServletRequest, UriInfo 등)는 요청 본문이 아니므로 무시한다.
	default:
		classifyBodyParam(typeName, paramName, ep)
	}
}
