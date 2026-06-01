//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what from-import 문에서 import된 이름 노드(name 필드)들을 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// fromImportNames returns the imported-name nodes of a from-import statement
// (the `name:` fields), each a dotted_name or aliased_import.
func fromImportNames(stmt *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	for i := 0; i < int(stmt.ChildCount()); i++ {
		if stmt.FieldNameForChild(i) == "name" {
			out = append(out, stmt.Child(i))
		}
	}
	return out
}
