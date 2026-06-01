//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 단일 decorated_definition에서 Flask 라우트 정보를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractOneRoute extracts route info from a single decorated definition.
// Returns nil if the definition is not a route handler.
// For @app.route with methods=["GET", "POST"], returns multiple routeInfo (one per method).
func extractOneRoute(def *sitter.Node, src []byte, bpPrefixes blueprintPrefix, file string) []routeInfo {
	decorators := childrenOfType(def, "decorator")
	if len(decorators) == 0 {
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

	bf := extractBodyFields(funcDef, src)

	var all []routeInfo
	for _, dec := range decorators {
		routes := parseFlaskDecorator(dec, src, bpPrefixes, handler, file, line)
		all = append(all, routes...)
	}
	if len(all) == 0 {
		return nil
	}

	return applyBodyFields(all, bf)
}
