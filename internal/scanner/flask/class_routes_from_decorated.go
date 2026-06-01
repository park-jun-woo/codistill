//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 단일 decorated_definition에서 @ns.route 클래스 라우트를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classRoutesFromDecorated extracts flask_restx routes from one decorated_definition.
// It returns nil unless the wrapped node is a Resource-subclass class_definition.
// Each *.route decorator on the class contributes one path (stacked @ns.route
// decorators yield a route per decorator); the namespace variable resolves the
// prefix via nsPrefixes, and the HTTP methods come from the class body.
func classRoutesFromDecorated(dd *sitter.Node, fi fileInfo, aliases importAlias, nsPrefixes namespacePrefix) []routeInfo {
	cls := findChildByType(dd, "class_definition")
	if cls == nil {
		return nil
	}
	if !isResourceSubclass(classSuperclasses(cls, fi.src), aliases) {
		return nil
	}
	nameNode := findChildByType(cls, "identifier")
	if nameNode == nil {
		return nil
	}
	className := nodeText(nameNode, fi.src)
	methods := classHTTPMethods(cls, fi.src)
	if len(methods) == 0 {
		return nil
	}

	var routes []routeInfo
	for _, dec := range childrenOfType(dd, "decorator") {
		nsVar, path, ok := nsRouteDecorator(dec, fi.src)
		if !ok {
			continue
		}
		prefix := nsPrefixes[nsVar]
		routes = append(routes, classRoutesToRouteInfo(methods, path, prefix, className+".", fi.relPath)...)
	}
	return routes
}
