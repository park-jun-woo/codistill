//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what argument_list에서 모든 positional 문자열 인자를 순서대로 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// allStringArgs collects every positional string argument in an argument list,
// in source order. Used by api.add_resource(Resource, path[, path2...]) where
// Flask-RESTful allows multiple alias paths for a single resource.
func allStringArgs(args *sitter.Node, src []byte) []string {
	var out []string
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "string" {
			out = append(out, unquotePython(nodeText(child, src)))
		}
	}
	return out
}
