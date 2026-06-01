//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what AST에서 이름이 일치하는 const 문자열 선언의 값을 찾는다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// lookupConstString scans root for a `const <name> = '<string>'` declaration
// (lexical_declaration > variable_declarator > identifier + string) and returns
// the unquoted string value. Non-string initializers (e.g. const X = foo())
// yield ("", false) so callers fall back to the raw identifier text.
func lookupConstString(root *sitter.Node, src []byte, name string) (string, bool) {
	for _, decl := range findAllByType(root, "variable_declarator") {
		nameNode := findChildByType(decl, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != name {
			continue
		}
		strNode := findChildByType(decl, "string")
		if strNode == nil {
			return "", false
		}
		return unquoteTS(nodeText(strNode, src)), true
	}
	return "", false
}
