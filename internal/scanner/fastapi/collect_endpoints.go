//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 모든 파일에서 엔드포인트와 모델 요청을 수집한다
package fastapi

import "github.com/park-jun-woo/codistill/internal/scanner"

// collectEndpoints extracts endpoints and model requests from all parsed files.
func collectEndpoints(files []fileInfo) ([]scanner.Endpoint, []modelRequest) {
	aliasMap := resolveTypeAliases(files)
	var endpoints []scanner.Endpoint
	var modelReqs []modelRequest
	for _, fi := range files {
		routes := extractRoutes(fi.root, fi.src, fi.prefixes, fi.routerDeps, fi.hidden, fi.relPath, aliasMap)
		for _, ri := range routes {
			epIdx := len(endpoints)
			endpoints = append(endpoints, buildEndpoint(ri))
			modelReqs = appendModelRequests(modelReqs, ri, fi, epIdx)
		}
	}
	return endpoints, modelReqs
}
