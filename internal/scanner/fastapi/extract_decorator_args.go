//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 데코레이터 인자에서 path, status_code, response_model, response_class, include_in_schema를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractDecoratorArgs extracts path, status_code, response_model, response_class,
// and include_in_schema from decorator arguments. include_in_schema defaults to
// true; it is false only when given as a literal False.
func extractDecoratorArgs(callNode *sitter.Node, src []byte) decoratorArgs {
	if callNode == nil {
		return decoratorArgs{includeInSchema: true}
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return decoratorArgs{includeInSchema: true}
	}

	return decoratorArgs{
		path:            firstStringArg(args, src),
		statusCode:      parseIntDefault(extractKeywordArg(args, "status_code", src), 0),
		responseModel:   extractKeywordArg(args, "response_model", src),
		responseClass:   extractKeywordArg(args, "response_class", src),
		includeInSchema: !keywordIsFalse(args, "include_in_schema", src),
	}
}
