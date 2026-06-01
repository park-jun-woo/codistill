//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 decorated_definition 에서 라우트 정보를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractOneRoute extracts route info from a single decorated definition.
// Returns nil if the definition is not a route handler, if the route decorator
// carries include_in_schema=False, or if its router variable is hidden.
// routerDeps maps router variable names to their constructor-level dependencies.
// hidden is the set of router variable names hidden from the schema.
// aliasMap maps type alias names to their Depends function names.
func extractOneRoute(def *sitter.Node, src []byte, prefixes map[string]string, routerDeps map[string][]string, hidden map[string]bool, file string, aliasMap map[string]string) *routeInfo {
	decorators := childrenOfType(def, "decorator")
	if len(decorators) == 0 {
		return nil
	}

	method, routerVar, da := findRouteDecorator(decorators, src)
	if method == "" {
		return nil
	}
	if !da.includeInSchema || hidden[routerVar] {
		return nil
	}

	funcDef := findChildByType(def, "function_definition")
	if funcDef == nil {
		return nil
	}

	nameNode := findChildByType(funcDef, "identifier")
	handler := ""
	if nameNode != nil {
		handler = nodeText(nameNode, src)
	}

	line := int(funcDef.StartPoint().Row) + 1
	prefix := prefixes[routerVar]
	fullPath := combinePath(prefix, da.path)

	ri := &routeInfo{
		method:        method,
		path:          fullPath,
		handler:       handler,
		file:          file,
		line:          line,
		statusCode:    da.statusCode,
		responseModel: da.responseModel,
		responseClass: da.responseClass,
	}

	if rDeps := routerDeps[routerVar]; len(rDeps) > 0 {
		ri.middleware = append(ri.middleware, rDeps...)
	}
	ri.middleware = append(ri.middleware, collectDecoratorDeps(decorators, src)...)

	extractParams(funcDef, src, ri, aliasMap)
	extractReturnType(funcDef, src, ri)

	return ri
}
