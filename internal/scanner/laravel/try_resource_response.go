//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what new SomeResource($var) 형태의 JSON resource 응답을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func tryResourceResponse(absRoot string, retNode *sitter.Node, src []byte, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *scanner.Response {
	for _, oc := range findAllByType(retNode, "object_creation_expression") {
		resName := resourceTypeName(oc, src)
		if resName == "" {
			continue
		}
		return &scanner.Response{
			Status:   "200",
			Kind:     "json",
			TypeName: resName,
			Fields:   extractResourceFields(absRoot, resName, srcFI, parsedFiles),
		}
	}
	return nil
}
