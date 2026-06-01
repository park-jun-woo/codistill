//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what register_blueprint 호출에서 (부모 변수, 자식 변수, prefix 오버라이드)를 파싱한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseRegisterBlueprintEdge parses a register_blueprint call into the nesting
// edge it expresses: which blueprint (parent receiver) registers which child,
// and any url_prefix override. e.g. `parent.register_blueprint(child, url_prefix="/v1")`
// -> ("parent", "child", "/v1"). Returns empty strings when the call is not a
// register_blueprint. The parent is the receiver before ".register_blueprint";
// app/api receivers (no blueprint prefix) simply resolve to no parent prefix.
func parseRegisterBlueprintEdge(call *sitter.Node, src []byte) (string, string, string) {
	attrNode := findChildByType(call, "attribute")
	if attrNode == nil {
		return "", "", ""
	}
	attrText := nodeText(attrNode, src)
	if !strings.HasSuffix(attrText, ".register_blueprint") {
		return "", "", ""
	}
	parent := strings.TrimSuffix(attrText, ".register_blueprint")
	if i := strings.LastIndex(parent, "."); i >= 0 {
		parent = parent[i+1:]
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return "", "", ""
	}
	childVar := firstIdentArg(args, src)
	if childVar == "" {
		return "", "", ""
	}
	override := extractKeywordArg(args, "url_prefix", src)
	return parent, childVar, override
}
