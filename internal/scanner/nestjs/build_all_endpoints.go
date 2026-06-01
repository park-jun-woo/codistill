//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 모든 컨트롤러에서 엔드포인트와 DTO 요청 목록을 생성한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildAllEndpoints iterates over all controllers and builds endpoint + DTO request lists.
func buildAllEndpoints(globalPrefix string, uriVersioning bool, defaultVersion string, controllers []controllerWithFile, projectRoot string) ([]scanner.Endpoint, []dtoRequest) {
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, cwf := range controllers {
		eps, reqs := buildControllerEndpoints(globalPrefix, uriVersioning, defaultVersion, cwf, projectRoot, len(endpoints))
		endpoints = append(endpoints, eps...)
		dtoReqs = append(dtoReqs, reqs...)
	}
	return endpoints, dtoReqs
}
