//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 call_expression에서 체인 또는 일반 라우트를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractRouteFromCall(call *sitter.Node, src []byte, routers map[string]bool, processed map[uintptr]bool) []routeInfo {
	chain := extractRouteChain(call, src, routers)
	if len(chain) > 0 {
		markChainProcessed(call, processed)
		return chain
	}
	routes := extractRoutesFromCall(call, src, routers)
	if len(routes) > 0 {
		processed[uintptr(call.StartByte())] = true
		return routes
	}
	return nil
}
