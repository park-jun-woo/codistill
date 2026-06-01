//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 어트리뷰트 없는 파라미터를 path param 또는 암묵(implicit) body로 분류한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func classifyImplicitParam(typeName, paramName string, ep *endpointInfo) {
	typeName, _ = stripNullable(typeName)
	if isPathParam(paramName, ep.path) {
		ep.params = append(ep.params, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
		return
	}
	if ep.bodyType != "" || isPrimitiveType(typeName) || !methodAllowsBody(ep.method) {
		return
	}
	ep.bodyType = typeName
	ep.bodyVarName = paramName
}
