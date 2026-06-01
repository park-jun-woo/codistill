//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what SomeResource::collection($var) 형태의 JSON 배열 응답을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func tryCollectionResponse(absRoot string, retNode *sitter.Node, src []byte, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *scanner.Response {
	for _, sc := range findAllByType(retNode, "scoped_call_expression") {
		name := collectionResourceName(sc, src)
		if name == "" {
			continue
		}
		return &scanner.Response{
			Status:   "200",
			Kind:     "json",
			TypeName: "[]" + name,
			Fields:   extractResourceFields(absRoot, name, srcFI, parsedFiles),
		}
	}
	return nil
}
