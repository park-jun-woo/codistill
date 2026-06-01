//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what object_pattern(구조분해)에 주어진 식별자명이 포함되는지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func objectPatternHasName(pattern *sitter.Node, src []byte, name string) bool {
	for _, sp := range childrenOfType(pattern, "shorthand_property_identifier_pattern") {
		if nodeText(sp, src) == name {
			return true
		}
	}
	return false
}
