//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what relative_import 노드에서 점 개수와 하위 dotted_name을 추출한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// relativeImportParts extracts the leading-dot count and optional sub dotted_name
// from a relative_import node (e.g. "from ..pkg import x" -> (2, "pkg")).
func relativeImportParts(mn *sitter.Node, src []byte) (int, string) {
	dots := 0
	sub := ""
	for i := 0; i < int(mn.ChildCount()); i++ {
		c := mn.Child(i)
		switch c.Type() {
		case "import_prefix":
			dots = strings.Count(nodeText(c, src), ".")
		case "dotted_name":
			sub = nodeText(c, src)
		}
	}
	return dots, sub
}
