//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 클래스/메서드 노드에 붙은 데코레이터 목록을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// findRouteDecorators returns decorators attached to a node.
// For exported classes, decorators are children of the parent export_statement.
// For methods, decorators are consecutive siblings preceding the method in the
// class_body.
func findRouteDecorators(node *sitter.Node, src []byte) []routeDecorator {
	parent := node.Parent()
	if parent == nil {
		return nil
	}
	if parent.Type() == "export_statement" {
		return collectDecoratorChildren(parent, src)
	}
	return collectPrecedingDecorators(parent, node, src)
}
