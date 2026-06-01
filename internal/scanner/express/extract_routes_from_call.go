//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 call_expression에서 (배열형 경로 전개 포함) HTTP 메서드 라우트를 다건 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractRoutesFromCall(call *sitter.Node, src []byte, routers map[string]bool) []routeInfo {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	routerVar := nodeText(obj, src)
	if !routers[routerVar] {
		return nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	upperMethod, ok := httpMethods[nodeText(prop, src)]
	if !ok {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	routes := buildRoutesFromArgs(args, src, upperMethod, int(call.StartPoint().Row)+1)
	for i := range routes {
		routes[i].Router = routerVar
	}
	return routes
}
