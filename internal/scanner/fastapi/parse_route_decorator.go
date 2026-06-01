//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what 데코레이터 노드에서 HTTP 메서드, 라우터 변수, 데코레이터 인자(path/status_code/response_model/response_class/include_in_schema)를 파싱한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// parseRouteDecorator parses a decorator like @app.get("/users/{user_id}", status_code=200).
// Returns (method, routerVar, args). args.includeInSchema defaults to true.
func parseRouteDecorator(dec *sitter.Node, src []byte) (string, string, decoratorArgs) {
	callNode, attrNode := findDecoratorNodes(dec)
	if callNode != nil && attrNode == nil {
		attrNode = findChildByType(callNode, "attribute")
	}
	if attrNode == nil {
		return "", "", decoratorArgs{includeInSchema: true}
	}

	routerVar, httpMethod := parseAttribute(attrNode, src)
	if httpMethod == "" {
		return "", "", decoratorArgs{includeInSchema: true}
	}

	return httpMethod, routerVar, extractDecoratorArgs(callNode, src)
}
