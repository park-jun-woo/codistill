//ff:func feature=scan type=extract control=selection topic=express
//ff:what this.route({...}) / this.<method>("...") 호출 한 건에서 method·path를 추출해 Endpoint를 만든다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractThisRouteCall interprets a single call_expression as a custom
// Controller route registration. Two forms are supported:
//
//  1. this.route({ method: 'get', path: '/tokens' }) — method/path are pulled
//     from the object literal's static string pairs.
//  2. this.get('/x', ...) / this.post(...) etc. — the member property is the
//     HTTP method and the first string argument is the path.
//
// The member object must be `this`; any other receiver yields false. handler is
// the enclosing method name. Calls missing a static method or path are skipped.
func extractThisRouteCall(call *sitter.Node, src []byte, handler, relPath string) (scanner.Endpoint, bool) {
	mem := findChildByType(call, "member_expression")
	if mem == nil || findChildByType(mem, "this") == nil {
		return scanner.Endpoint{}, false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return scanner.Endpoint{}, false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return scanner.Endpoint{}, false
	}
	propName := nodeText(prop, src)
	var method, path string
	switch propName {
	case "route":
		obj := findChildByType(args, "object")
		if obj == nil {
			return scanner.Endpoint{}, false
		}
		rawMethod := extractPairStringValue(obj, src, "method")
		upper, ok := httpMethods[rawMethod]
		if !ok {
			return scanner.Endpoint{}, false
		}
		method = upper
		path = extractPairStringValue(obj, src, "path")
	default:
		upper, ok := httpMethods[propName]
		if !ok {
			return scanner.Endpoint{}, false
		}
		method = upper
		strNode := findChildByType(args, "string")
		if strNode == nil {
			return scanner.Endpoint{}, false
		}
		path = unquoteTS(nodeText(strNode, src))
	}
	if path == "" {
		path = "/"
	}
	return scanner.Endpoint{
		Method:  method,
		Path:    expressPathToOpenAPI(path),
		Handler: handler,
		File:    relPath,
		Line:    int(call.StartPoint().Row) + 1,
	}, true
}
