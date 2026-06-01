//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 튜플 리스트 등록 호출의 첫 식별자 인자와 list 인자를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// configureAPIArgs inspects an argument_list for the (blueprint, [tuples]) shape:
// the first positional argument must be an identifier (the blueprint variable),
// and a list argument must follow. It returns the blueprint variable name and the
// list node. If the shape does not match, listNode is nil.
func configureAPIArgs(args *sitter.Node, src []byte) (bpVar string, listNode *sitter.Node) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "identifier" && bpVar == "" {
			bpVar = nodeText(child, src)
		}
		if child.Type() == "list" && bpVar != "" {
			return bpVar, child
		}
	}
	return "", nil
}
