//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what .route() 인자 노드에서 (method, handler, pathOverride)를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// parseRouteArg parses a .route() call's arguments. It returns the HTTP method
// and handler from the `web::method().to(h)` argument. When the first argument
// is a string literal (B-form: `App::new().route("/health", web::get().to(h))`),
// its content is returned as pathOverride so the caller registers the route at
// that path instead of the receiver-supplied resource path. A-form
// (`web::resource("/sub").route(web::get().to(h))`) has a call_expression first
// arg, not a string literal, so pathOverride is "" and behaviour is unchanged.
func parseRouteArg(args *sitter.Node, src []byte) (method, handler, pathOverride string) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.IsNamed() && child.Type() == "string_literal" {
			pathOverride = stringLiteralContent(child, src)
			continue
		}
		fe := routeToFieldExpr(child, src)
		if fe == nil {
			continue
		}
		method = extractWebMethod(fe, src)
		handler = extractToHandler(child, src)
	}
	return method, handler, pathOverride
}
