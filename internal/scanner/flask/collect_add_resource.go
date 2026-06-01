//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 AST에서 *.add_resource 호출을 찾아 등록 정보를 수집한다
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// collectAddResource scans a file AST for X.add_resource(Resource, path[, path2...])
// calls and returns one addResourceReg per matching call. The per-call parsing
// (receiver base, Resource class, URL rules) is delegated to parseAddResourceCall.
func collectAddResource(root *sitter.Node, src []byte) []addResourceReg {
	var regs []addResourceReg
	for _, call := range findAllByType(root, "call") {
		if reg, ok := parseAddResourceCall(call, src); ok {
			regs = append(regs, reg)
		}
	}
	return regs
}
