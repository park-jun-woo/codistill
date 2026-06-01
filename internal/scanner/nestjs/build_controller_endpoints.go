//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 컨트롤러에서 엔드포인트와 DTO 요청 목록을 생성한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildControllerEndpoints builds endpoints for a single controller.
// baseIdx is the current length of the global endpoints slice for correct DTO indexing.
func buildControllerEndpoints(globalPrefix string, uriVersioning bool, defaultVersion string, cwf controllerWithFile, projectRoot string, baseIdx int) ([]scanner.Endpoint, []dtoRequest) {
	ci := cwf.info
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ep := range ci.endpoints {
		// Array-path decorators (@Get(['/a','/b'])) fan out into one endpoint
		// per path. Single-path endpoints keep their original single path.
		pathList := ep.paths
		if len(pathList) <= 1 {
			pathList = []string{ep.path}
		}
		for _, p := range pathList {
			epCopy := ep
			epCopy.path = p
			endpoint := buildEndpoint(globalPrefix, uriVersioning, defaultVersion, ci, epCopy)
			// Recompute epIdx per cloned endpoint so DTO matching stays aligned.
			epIdx := baseIdx + len(endpoints)
			reqs := collectDTORequests(epCopy, ci.imports, cwf.absFile, projectRoot, epIdx)
			dtoReqs = append(dtoReqs, reqs...)
			endpoints = append(endpoints, endpoint)
		}
	}
	return endpoints, dtoReqs
}
