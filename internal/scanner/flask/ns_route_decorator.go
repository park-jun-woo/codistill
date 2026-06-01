//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what 클래스 데코레이터가 *.route(path)인지 판정해 ns 변수와 path를 반환한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// nsRouteDecorator inspects a class decorator and, when it is a *.route("/path")
// call (the flask_restx @ns.route form), returns the namespace variable (the
// attribute base, e.g. "inner_api_ns" in @inner_api_ns.route) and the first
// string argument as the raw route path. ok is false for non-route decorators
// or decorators without a string path argument.
func nsRouteDecorator(dec *sitter.Node, src []byte) (nsVar, path string, ok bool) {
	call := findChildByType(dec, "call")
	if call == nil {
		return "", "", false
	}
	attr := findChildByType(call, "attribute")
	if attr == nil {
		return "", "", false
	}
	attrText := nodeText(attr, src)
	parts := strings.SplitN(attrText, ".", 2)
	if len(parts) != 2 || parts[1] != "route" {
		return "", "", false
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return "", "", false
	}
	rawPath := firstStringArg(args, src)
	if rawPath == "" {
		return "", "", false
	}
	return parts[0], rawPath, true
}
