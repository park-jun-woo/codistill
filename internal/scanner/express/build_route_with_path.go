//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 주어진 path 문자열로 argNodes에서 핸들러·미들웨어·validator를 파싱해 routeInfo를 생성한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func buildRouteWithPath(argNodes []*sitter.Node, src []byte, method, path string, line int) *routeInfo {
	handler, middleware := extractHandlerAndMiddleware(argNodes, src)
	validators := extractZodValidatorsFromArgs(argNodes, src, 1)
	joiRefs := extractJoiRefsFromArgs(argNodes, src, 1)
	authLevel, roles := extractAuthFromArgs(argNodes, src)
	lastArg := argNodes[len(argNodes)-1]
	return &routeInfo{
		Method:        method,
		Path:          path,
		Handler:       handler,
		HandlerNode:   lastArg,
		Middleware:    middleware,
		Line:          line,
		ZodValidators: validators,
		JoiRefs:       joiRefs,
		AuthLevel:     authLevel,
		Roles:         roles,
	}
}
