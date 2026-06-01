//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 한 decorated 메서드의 @expose 데코레이터들에서 라우트를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// exposeRoutesFromMethod extracts routes from one decorated method of a
// Flask-AppBuilder API class. For each @expose decorator (parseExposeDecorator)
// it composes api.baseURL + path (OpenAPI form) and emits one routeInfo per HTTP
// method, labeled "ClassName.method". Non-expose decorators are skipped.
func exposeRoutesFromMethod(dd *sitter.Node, api appbuilderAPIInfo, aliases importAlias, src []byte) []routeInfo {
	fn := findChildByType(dd, "function_definition")
	if fn == nil {
		return nil
	}
	nameNode := findChildByType(fn, "identifier")
	if nameNode == nil {
		return nil
	}
	methodName := nodeText(nameNode, src)
	line := int(fn.StartPoint().Row) + 1
	var routes []routeInfo
	for _, dec := range childrenOfType(dd, "decorator") {
		rawPath, methods, ok := parseExposeDecorator(dec, src, aliases)
		if !ok {
			continue
		}
		routes = append(routes, exposeMethodRoutes(api, methodName, rawPath, methods, line)...)
	}
	return routes
}
