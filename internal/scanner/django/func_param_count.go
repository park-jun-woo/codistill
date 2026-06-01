//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 함수 정의의 파라미터 개수를 센다
package django

import sitter "github.com/smacker/go-tree-sitter"

// funcParamCount returns the number of formal parameters of a function_definition.
// It counts identifier, typed, default and starred parameter forms.
func funcParamCount(funcDef *sitter.Node, src []byte) int {
	params := findChildByType(funcDef, "parameters")
	if params == nil {
		return 0
	}
	count := 0
	for i := 0; i < int(params.ChildCount()); i++ {
		switch params.Child(i).Type() {
		case "identifier",
			"typed_parameter",
			"default_parameter",
			"typed_default_parameter",
			"list_splat_pattern",
			"dictionary_splat_pattern",
			"keyword_separator":
			count++
		}
	}
	return count
}
