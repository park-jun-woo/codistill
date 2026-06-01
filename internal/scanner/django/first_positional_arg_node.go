//ff:func feature=scan type=extract control=sequence topic=django
//ff:what call 노드의 argument_list에서 첫 위치 인자 노드를 반환한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// firstPositionalArgNode returns the first positional (non-keyword) argument
// node inside the call node's argument_list, or nil when there is none.
func firstPositionalArgNode(callNode *sitter.Node) *sitter.Node {
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	pos := positionalArgs(args)
	if len(pos) == 0 {
		return nil
	}
	return pos[0]
}
