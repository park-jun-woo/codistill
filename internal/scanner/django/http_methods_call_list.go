//ff:func feature=scan type=extract control=sequence topic=django
//ff:what @require_http_methods([...]) call 노드의 리스트 리터럴에서 HTTP 메서드 문자열을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// httpMethodsCallList returns the string methods listed in a
// @require_http_methods([...]) call. It returns nil when the call is not
// require_http_methods or has no list-literal argument.
func httpMethodsCallList(callNode *sitter.Node, src []byte) []string {
	funcNode := findChildByType(callNode, "identifier")
	if funcNode == nil || nodeText(funcNode, src) != "require_http_methods" {
		return nil
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	listNode := findChildByType(args, "list")
	if listNode == nil {
		return nil
	}
	return extractStringList(listNode, src)
}
