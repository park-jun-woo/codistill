//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 메서드 정의에서 HTTP 메서드 데코레이터를 찾아 Endpoint를 생성한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// decoratorMethodEndpoint inspects the decorators of a single method
// definition. If exactly one HTTP-method decorator (@Get/@Post/...) is present,
// it builds an endpoint whose path is the controller prefix joined with the
// decorator path argument (or just the prefix when the decorator has no string
// argument). Returns false for non-route methods.
func decoratorMethodEndpoint(m *sitter.Node, src []byte, prefix, relPath string) (scanner.Endpoint, bool) {
	for _, d := range findRouteDecorators(m, src) {
		method, ok := decoratorHTTPMethods[d.name]
		if !ok {
			continue
		}
		return buildDecoratorEndpoint(m, src, d, method, prefix, relPath), true
	}
	return scanner.Endpoint{}, false
}
