//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 할당문을 라우터 인스턴스화로 파싱 시도한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryParseRouterAssignment tries to parse a single assignment as a router instantiation.
func tryParseRouterAssignment(assign *sitter.Node, src []byte, routerSubclassNames map[string]bool) *routerInfo {
	left := findChildByType(assign, "identifier")
	if left == nil {
		return nil
	}
	call := findChildByType(assign, "call")
	if call == nil {
		return nil
	}
	funcNode := findChildByType(call, "identifier")
	if funcNode == nil {
		return nil
	}
	funcName := nodeText(funcNode, src)
	if !routerClassNames[funcName] && !routerSubclassNames[funcName] {
		return nil
	}

	ri := &routerInfo{
		varName:   nodeText(left, src),
		isFastAPI: funcName == "FastAPI",
	}
	args := findChildByType(call, "argument_list")
	if args != nil {
		ri.prefix = extractKeywordArg(args, "prefix", src)
		ri.middleware = findDependenciesKeyword(args, src)
		ri.hidden = keywordIsFalse(args, "include_in_schema", src)
	}
	return ri
}
