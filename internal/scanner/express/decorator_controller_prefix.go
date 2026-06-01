//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 클래스 노드에서 컨트롤러 데코레이터(@RestController/@Controller) prefix를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// decoratorControllerPrefix returns the path prefix from @RestController('/x')
// (or @Controller('/x')). The bool is false when the class has no such
// decorator, so it is not treated as a routing controller.
func decoratorControllerPrefix(cls *sitter.Node, src []byte) (string, bool) {
	for _, d := range findRouteDecorators(cls, src) {
		if controllerDecorators[d.name] {
			return d.arg, true
		}
	}
	return "", false
}
