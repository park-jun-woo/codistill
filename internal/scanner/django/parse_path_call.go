//ff:func feature=scan type=extract control=sequence topic=django
//ff:what path() 또는 re_path() 호출을 파싱하여 urlEntry를 반환한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parsePathCall parses a path("pattern", view) or path("pattern", include("module")) call.
func parsePathCall(callNode *sitter.Node, src []byte) *urlEntry {
	funcName := callFuncName(callNode, src)
	if funcName != "path" && funcName != "re_path" && funcName != "rest_path" {
		return nil
	}

	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}

	posArgs := positionalArgs(args)
	if len(posArgs) < 1 {
		return nil
	}

	pattern := ""
	if posArgs[0].Type() == "string" {
		pattern = unquotePython(nodeText(posArgs[0], src))
	}

	entry := &urlEntry{pattern: pattern}
	// rest_path is a method-keyword routing helper: rest_path("p", GET=v, POST=v2).
	// The HTTP methods arrive as keyword arguments rather than a single view arg.
	if funcName == "rest_path" {
		entry.methodViews = extractMethodKeywordViews(args, src)
	}
	if len(posArgs) >= 2 {
		resolveSecondArg(entry, posArgs[1], src)
	}
	return entry
}
