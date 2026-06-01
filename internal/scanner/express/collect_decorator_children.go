//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 부모 노드의 모든 데코레이터 자식을 수집한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// collectDecoratorChildren collects all decorator children of a parent node.
func collectDecoratorChildren(parent *sitter.Node, src []byte) []routeDecorator {
	var result []routeDecorator
	for _, dn := range childrenOfType(parent, "decorator") {
		d := parseRouteDecorator(dn, src)
		if d.name != "" {
			result = append(result, d)
		}
	}
	return result
}
