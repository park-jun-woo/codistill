//ff:func feature=scan type=parse control=iteration dimension=1 topic=express
//ff:what 데코레이터 노드에서 이름과 첫 문자열 인자(path)를 파싱한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// parseRouteDecorator extracts the decorator name and its first string
// argument. For @Post('/x', {options}) only the first string argument is taken
// as the path; option objects and other arguments are ignored. A bare
// identifier decorator (@Public) yields a name with hasArg=false.
func parseRouteDecorator(node *sitter.Node, src []byte) routeDecorator {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "identifier":
			return routeDecorator{name: nodeText(child, src)}
		case "call_expression":
			return parseRouteDecoratorCall(child, src)
		}
	}
	return routeDecorator{}
}
