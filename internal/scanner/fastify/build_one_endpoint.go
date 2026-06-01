//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 단일 Endpoint를 생성한다 (JSON Schema에서 request/response 추출 포함)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func buildOneEndpoint(method, oaPath string, r routeInfo, relPath string, pathParams []string, src []byte, vars map[string]*sitter.Node) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:  method,
		Path:    oaPath,
		Handler: r.Handler,
		File:    relPath,
		Line:    r.Line,
	}
	if r.Schema != nil {
		if si := extractJSONSchema(r.Schema, src); si != nil {
			ep.OperationID = si.OperationID
			ep.Summary = si.Summary
			ep.Description = si.Description
			ep.Tags = si.Tags
		}
	}
	req, hasReq := buildRequest(r, pathParams, src, vars)
	if hasReq {
		ep.Request = &req
	}
	ep.Responses = buildResponses(r, src)
	return ep
}
