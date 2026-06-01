//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what call_expression이 setGlobalPrefix인지 확인하고 접두사를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// trySetGlobalPrefix checks if a call_expression is setGlobalPrefix and returns
// the prefix. A direct string literal arg is read by firstStringArg; when that
// fails it resolves an identifier arg to a same-file const string
// (resolveSetPrefixIdentifier) or extracts the literal portion of a
// concatenation (extractBinaryLiteralPrefix), e.g. globalPrefix → 'api' or
// urlPrefix + 'api' → 'api'. astRoot is the parsed AST of the call's file,
// used for same-file const resolution. Non-resolvable args yield ("", false)
// so callers keep the Phase044 .env/config fallback.
func trySetGlobalPrefix(call *sitter.Node, astRoot *sitter.Node, src []byte) (string, bool) {
	memberAccess := findChildByType(call, "member_expression")
	if memberAccess == nil {
		return "", false
	}
	prop := findChildByType(memberAccess, "property_identifier")
	if prop == nil || nodeText(prop, src) != "setGlobalPrefix" {
		return "", false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", false
	}
	if prefix, ok := firstStringArg(args, src); ok {
		return prefix, true
	}
	if prefix, ok := resolveSetPrefixIdentifier(args, astRoot, src); ok {
		return prefix, true
	}
	return extractBinaryLiteralPrefix(args, astRoot, src)
}
