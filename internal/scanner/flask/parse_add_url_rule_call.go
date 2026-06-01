//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what 단일 call 노드를 X.add_url_rule(rule, endpoint, view, methods=...) 등록으로 파싱한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAddURLRuleCall parses one call node as an
// X.add_url_rule(rule, endpoint, view[, methods=(...)]) registration. The first
// positional string argument is the URL rule; the first positional identifier
// argument is the view (a RequestHandler class name in Indico, or a plain
// function in standard Flask). HTTP methods come from the methods= keyword
// (list or tuple); an empty result defaults to GET downstream. Indico's "!"
// rule prefix forces app-root resolution: a single leading "!" is stripped and
// recorded via appRoot, while any other "!" is preserved. The attribute base
// (the receiver before .add_url_rule) is recorded as blueprintVar. ok is false
// for non-add_url_rule calls or calls missing a URL rule.
func parseAddURLRuleCall(call *sitter.Node, src []byte) (addURLRuleReg, bool) {
	attr := findChildByType(call, "attribute")
	if attr == nil {
		return addURLRuleReg{}, false
	}
	attrText := nodeText(attr, src)
	if !strings.HasSuffix(attrText, ".add_url_rule") {
		return addURLRuleReg{}, false
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return addURLRuleReg{}, false
	}
	rawPath := firstStringArg(args, src)
	if rawPath == "" {
		return addURLRuleReg{}, false
	}
	appRoot := false
	if strings.HasPrefix(rawPath, "!") {
		appRoot = true
		rawPath = rawPath[1:]
	}
	handler := firstIdentArg(args, src)
	methods := extractMethodsArg(args, src)
	base := attrText[:len(attrText)-len(".add_url_rule")]
	return addURLRuleReg{
		rawPath:      rawPath,
		handler:      handler,
		methods:      methods,
		blueprintVar: base,
		appRoot:      appRoot,
		line:         int(call.StartPoint().Row) + 1,
	}, true
}
