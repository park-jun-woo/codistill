//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 커스텀 Controller 클래스의 각 메서드 본문에서 this 기반 라우트 등록을 추출한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// controllerClassEndpoints walks every method of a custom Controller subclass,
// collecting the this.route(...) / this.<method>(...) route registrations in
// each method body. The enclosing method name is used as the handler. The base
// Controller applies no path prefix on its own, so paths are taken verbatim;
// any mount prefix is supplied separately by the standard prefix propagation.
func controllerClassEndpoints(cls *sitter.Node, src []byte, relPath string) []scanner.Endpoint {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	var endpoints []scanner.Endpoint
	for _, m := range childrenOfType(body, "method_definition") {
		handler := methodDefinitionName(m, src)
		endpoints = append(endpoints, controllerMethodEndpoints(m, src, handler, relPath)...)
	}
	return endpoints
}
