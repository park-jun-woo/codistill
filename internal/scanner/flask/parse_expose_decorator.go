//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what 메서드 데코레이터가 @expose(path, methods=...)인지 판정해 path/methods를 반환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// parseExposeDecorator inspects a method decorator and, when it is a bare
// identifier call named "expose" (Flask-AppBuilder's @expose("/path",
// methods=[...]) form — note it is a plain identifier, not an attribute like
// @app.route), returns the first string argument as the raw path and the HTTP
// methods. Both list (methods=["POST"]) and tuple (methods=("POST",)) literals
// are accepted via extractMethodsArg; a missing methods= defaults to GET. ok is
// false for any non-expose decorator. exposeAlias maps a local import alias back
// to "expose" (collectImportAliases) so aliased imports are recognized.
func parseExposeDecorator(dec *sitter.Node, src []byte, aliases importAlias) (path string, methods []string, ok bool) {
	call := findChildByType(dec, "call")
	if call == nil {
		return "", nil, false
	}
	fn := call.ChildByFieldName("function")
	if fn == nil || fn.Type() != "identifier" {
		return "", nil, false
	}
	name := nodeText(fn, src)
	if name != "expose" && aliases[name] != "expose" {
		return "", nil, false
	}
	args := findChildByType(call, "argument_list")
	if args == nil {
		return "", nil, false
	}
	rawPath := firstStringArg(args, src)
	if rawPath == "" {
		return "", nil, false
	}
	methods = extractMethodsArg(args, src)
	if len(methods) == 0 {
		methods = []string{"GET"}
	}
	return rawPath, methods, true
}
