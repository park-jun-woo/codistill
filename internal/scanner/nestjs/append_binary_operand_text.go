//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what 연결식 피연산자 하나의 문자열/const 값을 누적 문자열에 더한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// appendBinaryOperandText appends the resolvable string text of a single
// binary_expression operand to *out. String literals are unquoted and appended
// verbatim; bare const-identifiers are resolved via same-file const strings
// (resolveOperandConstString); nested binary_expression operands recurse
// (collectBinaryLiteralOperands). Other operands contribute nothing. astRoot is
// the parsed AST of the call's file.
func appendBinaryOperandText(child *sitter.Node, astRoot *sitter.Node, src []byte, out *string) {
	switch child.Type() {
	case "binary_expression":
		collectBinaryLiteralOperands(child, astRoot, src, out)
	case "string", "template_string":
		*out += unquoteTS(nodeText(child, src))
	case "identifier":
		*out += resolveOperandConstString(nodeText(child, src), astRoot, src)
	}
}
