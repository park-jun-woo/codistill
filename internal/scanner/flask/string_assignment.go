//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what expression_statement이 `name = "string"` 대입이면 이름/값을 반환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// stringAssignment reports whether an expression_statement is a simple
// `identifier = "literal"` assignment and, if so, returns the left identifier
// name and the unquoted string value. ok is false for non-assignment statements
// or non-string right-hand sides. Used by classStringAttrs.
func stringAssignment(es *sitter.Node, src []byte) (name, value string, ok bool) {
	asg := findChildByType(es, "assignment")
	if asg == nil {
		return "", "", false
	}
	left := asg.ChildByFieldName("left")
	right := asg.ChildByFieldName("right")
	if left == nil || right == nil || right.Type() != "string" {
		return "", "", false
	}
	return nodeText(left, src), unquotePython(nodeText(right, src)), true
}
