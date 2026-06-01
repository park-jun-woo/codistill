//ff:func feature=scan type=parse control=iteration dimension=1 topic=flask
//ff:what 단일 Flask 데코레이터에서 라우트 메서드/경로를 파싱한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseFlaskDecorator parses a single Flask decorator and returns route info.
func parseFlaskDecorator(dec *sitter.Node, src []byte, bpPrefixes blueprintPrefix, handler, file string, line int) []routeInfo {
	// Find call node within decorator
	callNode := findChildByType(dec, "call")
	if callNode == nil {
		return nil
	}

	attrNode := findChildByType(callNode, "attribute")
	if attrNode == nil {
		return nil
	}

	attrText := nodeText(attrNode, src)
	parts := strings.SplitN(attrText, ".", 2)
	if len(parts) != 2 {
		return nil
	}

	routerVar := parts[0]
	methodName := parts[1]

	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}

	// Extract the path from the first string argument
	rawPath := firstStringArg(args, src)

	// Resolve blueprint prefix (file-scoped to disambiguate same-named "bp" vars)
	prefix := lookupBlueprintPrefix(bpPrefixes, file, routerVar)
	fullPath := combinePath(prefix, rawPath)

	// Extract URL params from the full path
	params := extractURLParams(fullPath)

	// Convert Flask path to OpenAPI format
	openAPIPath := flaskPathToOpenAPI(fullPath)

	// Determine HTTP methods
	methods := resolveHTTPMethods(methodName, args, src)
	if len(methods) == 0 {
		return nil
	}

	var routes []routeInfo
	for _, m := range methods {
		routes = append(routes, routeInfo{
			method:  m,
			path:    openAPIPath,
			handler: handler,
			file:    file,
			line:    line,
			params:  params,
		})
	}
	return routes
}
