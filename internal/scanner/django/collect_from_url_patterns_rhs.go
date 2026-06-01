//ff:func feature=scan type=extract control=sequence topic=django
//ff:what urlpatterns RHS가 list든 i18n_patterns(...) 같은 래퍼 call이든 path() 호출을 수집한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// collectFromURLPatternsRHS collects path() calls from a urlpatterns assignment RHS,
// handling both `urlpatterns = [...]` (list) and `urlpatterns = i18n_patterns(...)` (call wrapper).
func collectFromURLPatternsRHS(assignNode *sitter.Node, src []byte) []urlEntry {
	if listNode := findChildByType(assignNode, "list"); listNode != nil {
		return parsePathCallsInList(listNode, src)
	}
	if callNode := findChildByType(assignNode, "call"); callNode != nil {
		if argList := findChildByType(callNode, "argument_list"); argList != nil {
			return parsePathCallsInList(argList, src)
		}
	}
	if attrNode := findChildByType(assignNode, "attribute"); attrNode != nil {
		text := nodeText(attrNode, src)
		if strings.HasSuffix(text, ".urls") {
			rv := strings.TrimSuffix(text, ".urls")
			return []urlEntry{{isInclude: true, includeRouterVar: rv}}
		}
	}
	// Parenthesized / concatenated RHS: `urlpatterns = ( [path(...), ...] + extra )`.
	// The list literal is nested under parenthesized_expression > binary_operator, so the
	// direct-child lookups above miss it; recurse through those wrappers to reach the list.
	if parenNode := findChildByType(assignNode, "parenthesized_expression"); parenNode != nil {
		return collectPathCallsFromRHSNode(parenNode, src)
	}
	if binNode := findChildByType(assignNode, "binary_operator"); binNode != nil {
		return collectPathCallsFromRHSNode(binNode, src)
	}
	return nil
}
