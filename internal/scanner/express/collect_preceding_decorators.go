//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 지정 노드 앞에 연속된 데코레이터 형제 노드를 수집한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// collectPrecedingDecorators scans backwards from node within parent to find
// the consecutive run of decorators immediately preceding it. Comments between
// a decorator and its method are skipped (not treated as a boundary).
func collectPrecedingDecorators(parent, node *sitter.Node, src []byte) []routeDecorator {
	idx := nodeChildIndex(parent, node)
	var result []routeDecorator
	for i := idx - 1; i >= 0; i-- {
		child := parent.Child(i)
		if child.Type() == "comment" {
			continue
		}
		if child.Type() != "decorator" {
			break
		}
		d := parseRouteDecorator(child, src)
		if d.name != "" {
			result = append(result, d)
		}
	}
	return result
}
