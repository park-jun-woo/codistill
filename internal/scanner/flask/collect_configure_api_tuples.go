//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 튜플 리스트 형태의 라우트 등록 호출(blueprint, [(path, Resource)...])을 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectConfigureAPITuples scans a file AST for helper calls whose second
// positional argument is a list of (path, ResourceClass) tuples — the shape used
// by Zou's configure_api_from_blueprint(blueprint, route_tuples). The helper name
// is irrelevant; the structure (first arg identifier = blueprint var, second arg
// list of tuple(string, identifier)) is what is matched. Each tuple yields one
// addResourceReg with the blueprint variable carried as the prefix key.
func collectConfigureAPITuples(root *sitter.Node, src []byte) []addResourceReg {
	var regs []addResourceReg
	for _, call := range findAllByType(root, "call") {
		args := findChildByType(call, "argument_list")
		if args == nil {
			continue
		}
		bpVar, listNode := configureAPIArgs(args, src)
		if listNode == nil {
			continue
		}
		regs = append(regs, tupleListRegs(listNode, bpVar, src)...)
	}
	return regs
}
