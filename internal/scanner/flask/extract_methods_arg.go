//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 데코레이터 인자에서 methods=['GET', 'POST'] 목록을 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractMethodsArg extracts HTTP methods from the methods= keyword argument.
// e.g., methods=["GET", "POST"] -> ["GET", "POST"]. Both list (methods=["POST"])
// and tuple (methods=("POST",)) literals are accepted — Flask-AppBuilder's
// @expose uses the tuple form. Returns nil if no methods argument is found.
func extractMethodsArg(args *sitter.Node, src []byte) []string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil || nodeText(keyNode, src) != "methods" {
			continue
		}
		// Find the collection node: methods=["POST"] or methods=("POST",)
		collNode := findChildByType(child, "list")
		if collNode == nil {
			collNode = findChildByType(child, "tuple")
		}
		if collNode == nil {
			continue
		}
		return extractStringList(collNode, src)
	}
	return nil
}
