//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what 단일 call 노드를 X.add_resource(Resource, path...) 등록으로 파싱한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseAddResourceCall parses one call node as an X.add_resource(Resource, path...)
// registration. The first positional identifier argument is the Resource class;
// every positional string argument is a URL rule (Flask-RESTful permits multiple
// alias paths). The attribute base (the receiver before .add_resource, e.g. "api"
// or a flask_restx namespace variable) is recorded as blueprintVar for prefix
// resolution. ok is false for non-add_resource calls or calls missing a class/path.
func parseAddResourceCall(call *sitter.Node, src []byte) (addResourceReg, bool) {
	attr := findChildByType(call, "attribute")
	if attr == nil {
		return addResourceReg{}, false
	}
	attrText := nodeText(attr, src)
	if !strings.HasSuffix(attrText, ".add_resource") {
		return addResourceReg{}, false
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return addResourceReg{}, false
	}
	className := firstIdentArg(args, src)
	paths := allStringArgs(args, src)
	if className == "" || len(paths) == 0 {
		return addResourceReg{}, false
	}
	base := attrText[:len(attrText)-len(".add_resource")]
	return addResourceReg{className: className, paths: paths, blueprintVar: base}, true
}
