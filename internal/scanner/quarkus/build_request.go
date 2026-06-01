//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what endpointInfo로 scanner.Request를 생성한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildRequest(ep endpointInfo) *scanner.Request {
	req := &scanner.Request{
		PathParams: ep.params,
		Query:      ep.query,
		Headers:    ep.headers,
	}
	if len(ep.files) > 0 {
		req.Files = ep.files
	}
	if len(ep.formParams) > 0 {
		req.FormFields = ep.formParams
	}
	if ep.bodyType != "" {
		typeName := ep.bodyType
		if ep.bodyIsArray {
			// Re-apply the slice marker so the shared schema builder emits
			// type:array + items:$ref(X) via resolvePrimitiveSchema/bodySchema.
			typeName += "[]"
		}
		req.Body = &scanner.Body{
			VarName:  ep.bodyVarName,
			Method:   "JAXRSBody",
			TypeName: typeName,
		}
	}
	if !hasContent(req) {
		return nil
	}
	return req
}
