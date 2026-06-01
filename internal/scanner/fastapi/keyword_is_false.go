//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 키워드 인자가 리터럴 False 값을 가지는지 판별한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// keywordIsFalse reports whether the named keyword argument is present with a
// literal False value (e.g. include_in_schema=False). Non-literal/dynamic values
// (include_in_schema=settings.x) and True/omitted return false, so callers stay
// conservative and only act on a definite literal False.
func keywordIsFalse(args *sitter.Node, name string, src []byte) bool {
	return extractKeywordArg(args, name, src) == "False"
}
