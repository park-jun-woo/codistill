//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 노드가 App/scope 빌더에 직접 등록된 .route("<path>", ...) 최상위 호출인지 판별한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// isTopLevelRouteCall reports whether n is an App/scope builder's direct
// B-form `.route("<path>", web::method().to(h))` call — the 2-argument route
// whose first argument is a path string literal. Such a call registers a route
// directly on App/scope (no web::resource receiver carrying the path), so it
// acts as its own Pass2 entry point. A-form `.route(web::get().to(h))` has a
// call_expression first arg (not a string literal) and is excluded here; it is
// reached via its web::resource receiver chain instead.
func isTopLevelRouteCall(n *sitter.Node, src []byte) bool {
	if n.Type() != "call_expression" {
		return false
	}
	fe := findChildByType(n, "field_expression")
	if fe == nil {
		return false
	}
	fid := findChildByType(fe, "field_identifier")
	if fid == nil || nodeText(fid, src) != "route" {
		return false
	}
	args := findChildByType(n, "arguments")
	if args == nil {
		return false
	}
	first := firstArgExpr(args)
	return first != nil && first.Type() == "string_literal"
}
