//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 연결식 피연산자를 좌→우로 순회하며 리터럴/const 문자열을 누적한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectBinaryLiteralOperands walks a binary_expression in left-to-right
// source order, appending resolvable string text to *out via
// appendBinaryOperandText. Non-literal, non-resolvable operands contribute
// nothing, leaving a conservative literal prefix. astRoot is the parsed AST of
// the call's file (used for same-file const resolution).
func collectBinaryLiteralOperands(node *sitter.Node, astRoot *sitter.Node, src []byte, out *string) {
	for i := 0; i < int(node.ChildCount()); i++ {
		appendBinaryOperandText(node.Child(i), astRoot, src, out)
	}
}
