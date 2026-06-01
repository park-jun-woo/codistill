//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what root에서 `const <name> = '<string literal>'` 선언의 문자열 리터럴 값을 해소한다 (없으면 "")
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveConstString(root *sitter.Node, src []byte, name string) string {
	for _, decl := range findAllByType(root, "lexical_declaration") {
		if lit := constStringInDecl(decl, src, name); lit != "" {
			return lit
		}
	}
	return ""
}
