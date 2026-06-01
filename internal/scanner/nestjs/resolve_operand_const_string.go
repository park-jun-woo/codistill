//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 식별자 피연산자를 동일 파일 const 문자열로 해석해 반환한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// resolveOperandConstString resolves a bare identifier operand to a same-file
// `const name = '<string>'` value (lookupConstString). Returns "" when name is
// not a simple identifier or no matching const string exists, so unresolved
// operands contribute nothing to a concatenated prefix. astRoot is the parsed
// AST of the call's file.
func resolveOperandConstString(name string, astRoot *sitter.Node, src []byte) string {
	if !isSimpleConstIdent(name) {
		return ""
	}
	if v, ok := lookupConstString(astRoot, src, name); ok {
		return v
	}
	return ""
}
