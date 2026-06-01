//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 메서드 본문의 this 기반 라우트 호출들을 Endpoint로 추출한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// controllerMethodEndpoints collects every this.route(...) / this.<method>(...)
// route registration found inside a single method body, using handler as the
// handler name for each synthesized endpoint.
func controllerMethodEndpoints(m *sitter.Node, src []byte, handler, relPath string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, call := range findAllByType(m, "call_expression") {
		if ep, ok := extractThisRouteCall(call, src, handler, relPath); ok {
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints
}
