//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what setGlobalPrefix 연결식에서 문자열 리터럴 조각을 이어붙여 접두사를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractBinaryLiteralPrefix extracts a prefix from a setGlobalPrefix argument
// that is a string concatenation (e.g. setGlobalPrefix(urlPrefix + 'api')).
// It walks the binary_expression operands in source order, appending string
// literal fragments verbatim and resolving identifier operands via same-file
// const strings (lookupConstString); unresolved operands contribute nothing.
// astRoot is the parsed AST of the call's file. Returns ("", false) when the
// argument is not a binary_expression or yields no literal text, so callers
// keep the Phase044 .env/config fallback.
func extractBinaryLiteralPrefix(args *sitter.Node, astRoot *sitter.Node, src []byte) (string, bool) {
	bin := findChildByType(args, "binary_expression")
	if bin == nil {
		return "", false
	}
	var prefix string
	collectBinaryLiteralOperands(bin, astRoot, src, &prefix)
	if prefix == "" {
		return "", false
	}
	return prefix, true
}
