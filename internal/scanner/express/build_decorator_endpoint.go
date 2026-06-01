//ff:func feature=scan type=extract control=sequence topic=express
//ff:what HTTP 메서드 데코레이터 한 건으로부터 prefix와 결합한 Endpoint를 생성한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildDecoratorEndpoint builds an endpoint for a single HTTP-method decorator.
// The path is the controller prefix joined with the decorator's string argument
// (or just the prefix when the decorator has no argument), normalised to "/"
// when empty.
func buildDecoratorEndpoint(m *sitter.Node, src []byte, d routeDecorator, method, prefix, relPath string) scanner.Endpoint {
	methodPath := ""
	if d.hasArg {
		methodPath = d.arg
	}
	fullPath := joinExpressPath(prefix, methodPath)
	if fullPath == "" {
		fullPath = "/"
	}
	return scanner.Endpoint{
		Method:  method,
		Path:    expressPathToOpenAPI(fullPath),
		Handler: methodDefinitionName(m, src),
		File:    relPath,
		Line:    int(m.StartPoint().Row) + 1,
	}
}
