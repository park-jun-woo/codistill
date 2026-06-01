//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what class_definition 본문에서 HTTP 메서드 def를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classHTTPMethods scans a class_definition body and returns the HTTP method
// defs (get/post/put/delete/patch/options/head). Helper and dunder methods
// (e.g. __init__, __helper) are excluded. Used by class-based route scanners
// (Flask-RESTful / flask_restx / Flask-AppBuilder).
func classHTTPMethods(classNode *sitter.Node, src []byte) []classMethod {
	body := findChildByType(classNode, "block")
	if body == nil {
		return nil
	}
	var methods []classMethod
	for _, fn := range findAllByType(body, "function_definition") {
		nameNode := findChildByType(fn, "identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, src)
		http, ok := httpMethods[name]
		if !ok || http == "" {
			continue
		}
		methods = append(methods, classMethod{
			name:     http,
			line:     int(fn.StartPoint().Row) + 1,
			funcNode: fn,
		})
	}
	return methods
}
