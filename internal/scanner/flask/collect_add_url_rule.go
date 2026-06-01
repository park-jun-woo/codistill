//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 AST에서 *.add_url_rule 호출을 찾아 등록 정보를 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectAddURLRule scans a file AST for X.add_url_rule(rule, endpoint, view,
// methods=...) calls and returns one addURLRuleReg per matching call. The
// per-call parsing is delegated to parseAddURLRuleCall; the relative file path
// is stamped onto each registration here.
func collectAddURLRule(root *sitter.Node, src []byte, file string) []addURLRuleReg {
	var regs []addURLRuleReg
	for _, call := range findAllByType(root, "call") {
		if reg, ok := parseAddURLRuleCall(call, src); ok {
			reg.file = file
			regs = append(regs, reg)
		}
	}
	return regs
}
