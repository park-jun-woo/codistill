//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what use 절의 별칭(use X\Y as Z의 Z)을 반환한다(없으면 빈 문자열)
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// useAlias returns the alias of a `use X\Y as Z` namespace_use_clause, or ""
// when the clause has no aliasing_clause.
func useAlias(use *sitter.Node, src []byte) string {
	clause := findChildByType(use, "namespace_aliasing_clause")
	if clause == nil {
		return ""
	}
	name := findChildByType(clause, "name")
	if name == nil {
		return ""
	}
	return nodeText(name, src)
}
