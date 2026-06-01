//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터 목록에서 첫 번째 HTTP 라우트 데코레이터를 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findRouteDecorator iterates decorators and returns the first HTTP route match
// as (method, routerVar, args).
func findRouteDecorator(decorators []*sitter.Node, src []byte) (string, string, decoratorArgs) {
	for _, dec := range decorators {
		m, rv, da := parseRouteDecorator(dec, src)
		if m != "" {
			return m, rv, da
		}
	}
	return "", "", decoratorArgs{includeInSchema: true}
}
