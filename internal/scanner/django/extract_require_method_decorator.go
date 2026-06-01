//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what plain 함수 뷰의 @require_POST/@require_GET/@require_http_methods 데코레이터에서 HTTP 메서드를 추출한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// requireMethodDecoratorNames maps Django's django.views.decorators.http
// shortcut decorators to the HTTP method(s) they whitelist.
var requireMethodDecoratorNames = map[string][]string{
	"require_POST":   {"POST"},
	"require_GET":    {"GET"},
	"require_safe":   {"GET", "HEAD"},
	"require_DELETE": {"DELETE"},
}

// extractRequireMethodDecorator inspects a function definition's decorators for
// Django HTTP-method-restricting decorators and returns the allowed HTTP methods.
//
//   - @require_POST / @require_GET / @require_safe / @require_DELETE -> fixed method(s)
//   - @require_http_methods(["POST", "PUT"])                         -> the listed methods
//
// Wrapper decorators (e.g. @csrf_exempt) and unrelated decorators are ignored, so
// they may freely coexist with a method-restricting decorator. Returns nil when no
// method-restricting decorator is present (callers keep the GET fallback).
func extractRequireMethodDecorator(funcDef *sitter.Node, src []byte) []string {
	parent := funcDef.Parent()
	if parent == nil || parent.Type() != "decorated_definition" {
		return nil
	}
	var raw []string
	for _, dec := range childrenOfType(parent, "decorator") {
		raw = append(raw, methodsFromDecorator(dec, src)...)
	}
	methods := dedupUpperMethods(raw)
	if len(methods) == 0 {
		return nil
	}
	return methods
}
