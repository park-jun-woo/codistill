//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import_from_statement의 aliased_import들을 importInfo 목록으로 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractAliasImports collects aliased imports (`import X as Y`) from an
// import_from_statement as importInfo entries with name=alias and origName=X.
func extractAliasImports(stmt *sitter.Node, module string, src []byte) []importInfo {
	var out []importInfo
	for _, an := range findAllByType(stmt, "aliased_import") {
		alias, orig := extractAliasImport(an, src)
		if alias == "" {
			continue
		}
		out = append(out, importInfo{name: alias, module: module, origName: orig})
	}
	return out
}
