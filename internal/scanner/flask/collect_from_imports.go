//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일의 from-import 문에서 로컬명 → 원본모듈/원본명 바인딩을 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectFromImports walks `from <module> import <name> [as <local>]` statements
// in a file and maps each locally bound name to the module it came from. This
// lets a route's blueprint variable (e.g. `bp` referenced by @bp.route after
// `from controllers.trigger import bp`) be resolved to the blueprint defined in
// that exact module, instead of whichever same-named blueprint was collected last.
func collectFromImports(root *sitter.Node, src []byte, relPath string) map[string]fromImportBinding {
	bindings := make(map[string]fromImportBinding)
	pkg := packageOf(relPath)
	for _, stmt := range findAllByType(root, "import_from_statement") {
		module := fromImportModule(stmt, src, pkg)
		if module == "" {
			continue
		}
		addStmtBindings(stmt, src, module, bindings)
	}
	return bindings
}
