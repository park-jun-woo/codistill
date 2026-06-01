//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 단일 return 문에서 응답 정보를 추출한다(Resource/Collection/json/noContent)
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractOneResponse extracts response info from a single return statement.
func extractOneResponse(absRoot string, retNode *sitter.Node, src []byte, srcFI *fileInfo, parsedFiles map[string]*fileInfo) *scanner.Response {
	if resp := tryResourceResponse(absRoot, retNode, src, srcFI, parsedFiles); resp != nil {
		return resp
	}
	if resp := tryCollectionResponse(absRoot, retNode, src, srcFI, parsedFiles); resp != nil {
		return resp
	}
	text := nodeText(retNode, src)
	if resp := tryJSONResponse(retNode, src, text); resp != nil {
		return resp
	}
	return tryNoContentResponse(text)
}
