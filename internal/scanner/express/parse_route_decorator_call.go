//ff:func feature=scan type=parse control=iteration dimension=1 topic=express
//ff:what 데코레이터 호출식(@Name('path', {opts}))에서 이름과 첫 문자열 path 인자를 파싱한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// parseRouteDecoratorCall handles @Name('path') or @Name('path', {opts}).
// Only the first string-literal argument is taken as the route path; option
// objects and non-string arguments are skipped so a leading options object or
// trailing options object never pollutes the path.
func parseRouteDecoratorCall(call *sitter.Node, src []byte) routeDecorator {
	d := routeDecorator{}
	fn := findChildByType(call, "identifier")
	if fn != nil {
		d.name = nodeText(fn, src)
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return d
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		arg := args.Child(i)
		switch arg.Type() {
		case "string", "template_string":
			d.arg = unquoteTS(nodeText(arg, src))
			d.hasArg = true
			return d
		}
	}
	return d
}
