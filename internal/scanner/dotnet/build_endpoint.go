//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what controllerInfo와 endpointInfo로 scanner.Endpoint를 생성한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpoint(ci controllerInfo, ep endpointInfo) scanner.Endpoint {
	actionPath := expandRouteTokens(ep.path, ci.className, ep.handler)
	var fullPath string
	if stripped, absolute := isAbsoluteRoute(actionPath); absolute {
		fullPath = joinPath(stripped)
	} else {
		fullPath = joinPath(ci.prefix, actionPath)
	}

	roles := mergeRoles(ci.roles, ep.roles)

	endpoint := scanner.Endpoint{
		Method:  ep.method,
		Path:    fullPath,
		Handler: ep.handler,
		File:    ep.file,
		Line:    ep.line,
		Roles:   roles,
	}

	req := buildRequest(ep)
	if req != nil {
		endpoint.Request = req
	}
	resp := buildResponse(ep)
	if resp != nil {
		endpoint.Responses = []scanner.Response{*resp}
	}
	return endpoint
}
