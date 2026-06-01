//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 클래스가 베이스 Controller를 extends 하는지 검사한다(커스텀 컨트롤러 게이트)
package express

import sitter "github.com/smacker/go-tree-sitter"

// classExtendsController reports whether the class declaration extends a base
// class named exactly "Controller" (Unleash-style custom controller base).
// This is the recognition gate for the custom-controller pass: only such
// classes have their this.route(...) / this.<method>(...) calls interpreted as
// routes, avoiding false positives from arbitrary classes that call this.get.
func classExtendsController(cls *sitter.Node, src []byte) bool {
	heritage := findChildByType(cls, "class_heritage")
	if heritage == nil {
		return false
	}
	for _, ext := range childrenOfType(heritage, "extends_clause") {
		id := findChildByType(ext, "identifier")
		if id != nil && nodeText(id, src) == "Controller" {
			return true
		}
	}
	return false
}
