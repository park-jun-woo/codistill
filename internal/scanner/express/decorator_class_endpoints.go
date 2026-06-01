//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 컨트롤러 클래스 본문의 메서드들에서 데코레이터 라우트 Endpoint를 생성한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// decoratorClassEndpoints builds endpoints from the @Get/@Post/... decorated
// methods of a single controller class, joining each method path onto prefix.
func decoratorClassEndpoints(cls *sitter.Node, src []byte, prefix, relPath string) []scanner.Endpoint {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	var endpoints []scanner.Endpoint
	for _, m := range childrenOfType(body, "method_definition") {
		ep, ok := decoratorMethodEndpoint(m, src, prefix, relPath)
		if ok {
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints
}
