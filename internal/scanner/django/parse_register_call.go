//ff:func feature=scan type=extract control=sequence topic=django
//ff:what router.register() 호출 노드를 파싱한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseRegisterCall parses a router.register() call node.
func parseRegisterCall(callNode *sitter.Node, fi fileInfo) *routerRegistration {
	attrNode := findChildByType(callNode, "attribute")
	if attrNode == nil {
		return nil
	}
	text := nodeText(attrNode, fi.src)
	routerVar, ok := registerRouterVar(text)
	if !ok {
		return nil
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	posArgs := positionalArgs(args)
	if len(posArgs) < 2 {
		return nil
	}
	prefix := ""
	if posArgs[0].Type() == "string" {
		prefix = unquotePython(nodeText(posArgs[0], fi.src))
	}
	return &routerRegistration{
		prefix:      strings.TrimRight(prefix, "/"),
		viewsetName: nodeText(posArgs[1], fi.src),
		basename:    extractKeywordArg(args, "basename", fi.src),
		routerVar:   routerVar,
		module:      fi.module,
	}
}
