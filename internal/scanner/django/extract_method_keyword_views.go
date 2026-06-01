//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what argument_list에서 HTTP 메서드명 키워드 인자(GET=view)를 메서드→뷰 맵으로 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractMethodKeywordViews scans an argument_list for keyword arguments whose
// names are HTTP methods (GET, POST, PUT, PATCH, DELETE, ...) and whose values
// are view references. It returns an uppercase HTTP-method -> view-name map.
// This models method-keyword routing helpers like
// rest_path("messages", GET=get_messages, POST=send_message). Returns nil when
// no HTTP-method keyword argument is present.
func extractMethodKeywordViews(args *sitter.Node, src []byte) map[string]string {
	result := map[string]string{}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil {
			continue
		}
		method, ok := apiviewHTTPMethods[strings.ToLower(nodeText(keyNode, src))]
		if !ok {
			continue
		}
		view := keywordArgViewValue(child, keyNode, src)
		if view != "" {
			result[method] = view
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
