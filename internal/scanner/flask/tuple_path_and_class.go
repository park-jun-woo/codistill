//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what (path, ResourceClass) 튜플에서 path 문자열과 클래스 식별자를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// tuplePathAndClass extracts the path string and class identifier from a
// (path, ResourceClass) tuple node, scanning only the tuple's direct children so
// nested expressions are ignored. The first string element is the path; the first
// identifier element is the Resource class.
func tuplePathAndClass(tup *sitter.Node, src []byte) (path, className string) {
	for i := 0; i < int(tup.ChildCount()); i++ {
		child := tup.Child(i)
		if child.Type() == "string" && path == "" {
			path = unquotePython(nodeText(child, src))
		}
		if child.Type() == "identifier" && className == "" {
			className = nodeText(child, src)
		}
	}
	return path, className
}
