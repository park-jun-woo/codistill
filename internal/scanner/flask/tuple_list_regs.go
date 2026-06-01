//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what (path, Resource) 튜플 리스트 노드를 등록 정보 슬라이스로 변환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// tupleListRegs converts a list node of (path, ResourceClass) tuples into
// addResourceReg entries, carrying the blueprint variable as the prefix key.
// Tuples missing a path or class are skipped.
func tupleListRegs(listNode *sitter.Node, bpVar string, src []byte) []addResourceReg {
	var regs []addResourceReg
	for _, tup := range findAllByType(listNode, "tuple") {
		path, className := tuplePathAndClass(tup, src)
		if path == "" || className == "" {
			continue
		}
		regs = append(regs, addResourceReg{
			className:    className,
			paths:        []string{path},
			blueprintVar: bpVar,
		})
	}
	return regs
}
