//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 데코레이터 컨트롤러 클래스에서 메서드 데코레이터 라우트를 추출하여 Endpoint를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

// extractDecoratorEndpoints scans a single file's class declarations for
// @RestController/@Controller-decorated controllers and builds an endpoint per
// @Get/@Post/@Put/@Patch/@Delete-decorated method, synthesizing the full path
// from the controller prefix and the method path.
func extractDecoratorEndpoints(fi *fileInfo, relPath string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	classes := findAllByType(fi.Root, "class_declaration")
	for _, cls := range classes {
		prefix, ok := decoratorControllerPrefix(cls, fi.Src)
		if !ok {
			continue
		}
		endpoints = append(endpoints, decoratorClassEndpoints(cls, fi.Src, prefix, relPath)...)
	}
	return endpoints
}
