//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 단일 from-import 문의 import된 이름들을 바인딩 맵에 추가한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// addStmtBindings records each imported name of one from-import statement into
// the bindings map, mapping the local name to its source module and original name.
func addStmtBindings(stmt *sitter.Node, src []byte, module string, bindings map[string]fromImportBinding) {
	for _, n := range fromImportNames(stmt) {
		local, orig := importedName(n, src)
		if local == "" {
			continue
		}
		bindings[local] = fromImportBinding{module: module, orig: orig}
	}
}
