//ff:func feature=scan type=extract control=sequence topic=express
//ff:what export_statement 한 건이 VERB export면 파일기반 Endpoint를 생성한다 (아니면 ok=false)
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildFilebasedEndpoint builds a single file-based endpoint from one
// export_statement. It returns ok=false when the export is not an uppercase
// HTTP verb (e.g. `export const config = {}`). The handler body is passed
// through extractResponses via a synthesized routeInfo so response extraction
// matches the raw express path.
func buildFilebasedEndpoint(exp *sitter.Node, fi *fileInfo, oaPath, relPath string) (scanner.Endpoint, bool) {
	method, handler, ok := filebasedVerbExport(exp, fi.Src)
	if !ok {
		return scanner.Endpoint{}, false
	}
	line := int(exp.StartPoint().Row) + 1
	ri := routeInfo{
		Method:      method,
		Path:        oaPath,
		Handler:     method,
		HandlerNode: handler,
		Line:        line,
	}
	return scanner.Endpoint{
		Method:    method,
		Path:      oaPath,
		Handler:   method,
		File:      relPath,
		Line:      line,
		Responses: extractResponses(fi, ri),
	}, true
}
