//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 인라인 require('...') / import('...') call_expression 인자에서 마운트 대상 파일 경로를 해소한다 (실패 시 "")
package express

import (
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

func resolveInlineRequirePath(callNode *sitter.Node, src []byte, fi *fileInfo, absRoot string, aliases map[string]string) string {
	modPath := extractRequirePath(callNode, src)
	if modPath == "" {
		return ""
	}
	dir := filepath.Dir(fi.Path)
	resolved := resolveRelativePath(dir, modPath)
	if resolved == "" {
		resolved = resolvePathAlias(absRoot, modPath, aliases)
	}
	return resolved
}
