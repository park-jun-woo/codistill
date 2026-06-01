//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what endpointInfo와 컨트롤러 정보로 scanner.Endpoint를 생성한다
package nestjs

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildEndpoint creates a scanner.Endpoint from controller + endpoint info.
func buildEndpoint(globalPrefix string, uriVersioning bool, defaultVersion string, ci controllerInfo, ep endpointInfo) scanner.Endpoint {
	versionPrefix := ""
	// Controller-level version wins; otherwise fall back to the app-level
	// defaultVersion from enableVersioning({ defaultVersion }).
	effectiveVersion := ci.version
	if effectiveVersion == "" {
		effectiveVersion = defaultVersion
	}
	if uriVersioning && effectiveVersion != "" {
		versionPrefix = "v" + effectiveVersion
	}
	fullPath := joinParts(globalPrefix, versionPrefix, ci.prefix, ep.path)
	fullPath = pathToOpenAPI(fullPath)
	if !strings.HasPrefix(fullPath, "/") {
		fullPath = "/" + fullPath
	}
	endpoint := scanner.Endpoint{
		Method:     ep.method,
		Path:       fullPath,
		Handler:    ep.handler,
		File:       ep.file,
		Line:       ep.line,
		Middleware: ep.middleware,
		Roles:      ep.roles,
		AuthLevel:  ep.authLevel,
	}
	req := buildRequest(ep)
	if req != nil {
		endpoint.Request = req
	}
	if ep.statusCode > 0 || ep.returnType != "" {
		endpoint.Responses = []scanner.Response{buildResponse(ep)}
	}
	return endpoint
}
