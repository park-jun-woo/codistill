//ff:func feature=scan type=extract control=selection topic=dotnet
//ff:what 파라미터 어트리뷰트에 따라 path/query/body/file/form/header로 분류한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyParam(param *sitter.Node, src []byte, ep *endpointInfo) {
	typeName := extractParamType(param, src)
	paramName := extractParamName(param, src)
	defaultVal := extractParamDefault(param, src)

	if isDIType(typeName) {
		return
	}

	switch {
	case hasAttribute(param, src, AttrFromBody):
		base, _ := stripNullable(typeName)
		ep.bodyType = base
		ep.bodyVarName = paramName
	case hasAttribute(param, src, AttrFromQuery):
		p := scanner.Param{Name: paramName, Type: csharpTypeToOpenAPIType(typeName)}
		if defaultVal != "" {
			p.Default = defaultVal
		}
		ep.query = append(ep.query, p)
	case hasAttribute(param, src, AttrFromRoute):
		ep.params = append(ep.params, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	case hasAttribute(param, src, AttrFromHeader):
		ep.headers = append(ep.headers, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	case hasAttribute(param, src, AttrFromForm):
		classifyFormParam(ep, typeName, paramName)
	default:
		classifyImplicitParam(typeName, paramName, ep)
	}
}
