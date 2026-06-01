//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일 내 Controller 서브클래스에서 this 기반 라우트 Endpoint를 추출한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

// extractControllerEndpoints scans a single file's class declarations for
// classes that extend the base Controller and builds endpoints from their
// this.route(...) / this.<method>(...) registrations. Classes that do not
// extend Controller contribute nothing (regression-safe gate).
func extractControllerEndpoints(fi *fileInfo, relPath string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, cls := range findAllByType(fi.Root, "class_declaration") {
		if !classExtendsController(cls, fi.Src) {
			continue
		}
		endpoints = append(endpoints, controllerClassEndpoints(cls, fi.Src, relPath)...)
	}
	return endpoints
}
