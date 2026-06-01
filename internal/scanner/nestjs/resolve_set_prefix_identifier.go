//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what setGlobalPrefix 인자가 식별자면 동일 파일 const 문자열로 해석한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// resolveSetPrefixIdentifier resolves a setGlobalPrefix argument that is a bare
// identifier (e.g. setGlobalPrefix(globalPrefix)) by looking up a same-file
// `const globalPrefix = '<string>'` declaration. astRoot is the parsed AST of
// the file containing the call. Returns ("", false) when the argument is not a
// simple identifier or no matching const string is found, so callers can fall
// back to the binary-expression resolver or .env/config fallback (Phase044).
func resolveSetPrefixIdentifier(args *sitter.Node, astRoot *sitter.Node, src []byte) (string, bool) {
	ident := findChildByType(args, "identifier")
	if ident == nil {
		return "", false
	}
	name := nodeText(ident, src)
	if !isSimpleConstIdent(name) {
		return "", false
	}
	return lookupConstString(astRoot, src, name)
}
