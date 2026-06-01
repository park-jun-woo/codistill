//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 함수 정의 본문에 *.register(...) 호출이 있는지 판정한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// funcBodyCallsRegister reports whether a function_definition body contains a
// call to a known register method (e.g. router.register(...)). This is the
// conservative gate for promoting a module-local helper to a router-register
// wrapper (PostHog's register_grandfathered_* family).
func funcBodyCallsRegister(funcDef *sitter.Node, src []byte) bool {
	for _, callNode := range findAllByType(funcDef, "call") {
		attrNode := findChildByType(callNode, "attribute")
		if attrNode == nil {
			continue
		}
		if _, ok := registerRouterVar(nodeText(attrNode, src)); ok {
			return true
		}
	}
	return false
}
