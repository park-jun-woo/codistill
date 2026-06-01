//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 현재 파일에서 `const {name} = require('...')` 구조분해 import의 모듈 파일 경로를 해소한다 (없으면 "")
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveDestructuredRequirePath(root *sitter.Node, src []byte, name, dir, absRoot string, aliases map[string]string) string {
	for _, decl := range findAllByType(root, "lexical_declaration") {
		if p := destructuredRequirePathInDecl(decl, src, name, dir, absRoot, aliases); p != "" {
			return p
		}
	}
	return ""
}
