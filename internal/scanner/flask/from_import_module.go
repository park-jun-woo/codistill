//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what from-import의 module_name을 절대 모듈 경로로 해석한다(상대 import 확장 포함)
package flask

import sitter "github.com/smacker/go-tree-sitter"

// fromImportModule resolves the module_name of a from-import to an absolute
// dotted path, expanding relative imports against the importing file's package.
func fromImportModule(stmt *sitter.Node, src []byte, pkg string) string {
	mn := stmt.ChildByFieldName("module_name")
	if mn == nil {
		return ""
	}
	if mn.Type() == "dotted_name" {
		return nodeText(mn, src)
	}
	if mn.Type() != "relative_import" {
		return ""
	}
	dots, sub := relativeImportParts(mn, src)
	if dots == 0 {
		return ""
	}
	return resolveRelativeModule(pkg, dots, sub)
}
