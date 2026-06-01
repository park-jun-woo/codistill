//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 lexical_declaration에서 `const {name} = require('...')` 구조분해의 모듈 파일 경로를 해소한다 (없으면 "")
package express

import (
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

func destructuredRequirePathInDecl(decl *sitter.Node, src []byte, name, dir, absRoot string, aliases map[string]string) string {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		pattern := findChildByType(declarator, "object_pattern")
		if pattern == nil || !objectPatternHasName(pattern, src, name) {
			continue
		}
		modPath := requirePathOfDeclarator(declarator, src)
		if modPath == "" {
			continue
		}
		resolved := resolveRelativePath(filepath.Clean(dir), modPath)
		if resolved == "" {
			resolved = resolvePathAlias(absRoot, modPath, aliases)
		}
		return resolved
	}
	return ""
}
