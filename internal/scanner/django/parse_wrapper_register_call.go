//ff:func feature=scan type=extract control=sequence topic=django
//ff:what register wrapper 함수 호출을 routerRegistration으로 파싱한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseWrapperRegisterCall parses a call to a known register-wrapper helper
// (e.g. register_grandfathered_environment_nested_viewset(r"insights", ViewSet,
// ...)). The first positional argument is the prefix and the second is the
// viewset, mirroring router.register("prefix", ViewSet). Returns nil if the
// call is not a known wrapper or its arguments do not fit the (str, viewset)
// shape.
func parseWrapperRegisterCall(callNode *sitter.Node, fi fileInfo, wrappers map[string]bool) *routerRegistration {
	idNode := findChildByType(callNode, "identifier")
	if idNode == nil {
		return nil
	}
	if !wrappers[nodeText(idNode, fi.src)] {
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
	if posArgs[0].Type() != "string" {
		return nil
	}
	prefix := unquotePython(nodeText(posArgs[0], fi.src))
	return &routerRegistration{
		prefix:      strings.TrimRight(prefix, "/"),
		viewsetName: nodeText(posArgs[1], fi.src),
		basename:    extractKeywordArg(args, "basename", fi.src),
		module:      fi.module,
	}
}
