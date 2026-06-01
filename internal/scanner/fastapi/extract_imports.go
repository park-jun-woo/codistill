//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what AST에서 import_from_statement 를 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractImports collects import information from the AST.
func extractImports(root *sitter.Node, src []byte) []importInfo {
	var imports []importInfo
	stmts := findAllByType(root, "import_from_statement")
	for _, stmt := range stmts {
		module := extractImportModule(stmt, src)
		if module == "" {
			continue
		}
		names := extractImportNames(stmt, src)
		for _, name := range names {
			imports = append(imports, importInfo{name: name, module: module, origName: name})
		}
		imports = append(imports, extractAliasImports(stmt, module, src)...)
	}
	return imports
}
