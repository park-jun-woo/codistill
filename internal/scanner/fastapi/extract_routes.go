//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터에서 HTTP 라우트를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractRoutes finds all decorated route functions in the AST and returns routeInfo slices.
// prefixes maps router variable names to their resolved path prefixes.
// routerDeps maps router variable names to their constructor-level dependencies.
// hidden is the set of router variable names hidden from the schema.
// aliasMap maps type alias names to their Depends function names.
func extractRoutes(root *sitter.Node, src []byte, prefixes map[string]string, routerDeps map[string][]string, hidden map[string]bool, file string, aliasMap map[string]string) []routeInfo {
	var routes []routeInfo
	defs := findAllByType(root, "decorated_definition")
	for _, def := range defs {
		ri := extractOneRoute(def, src, prefixes, routerDeps, hidden, file, aliasMap)
		if ri != nil {
			routes = append(routes, *ri)
		}
	}
	return routes
}
