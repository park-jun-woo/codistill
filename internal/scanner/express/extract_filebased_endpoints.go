//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 Medusa route.ts 파일에서 export const VERB 핸들러마다 파일기반 Endpoint를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

// extractFilebasedEndpoints builds one endpoint per `export const <VERB>`
// (or `export async function <VERB>`) handler found in a Medusa file-based
// route file. The path is derived from the file's directory structure
// (filebasedRoutePath). Handler bodies are passed through the standard response
// extraction (extractResponses) via a synthesized routeInfo.
func extractFilebasedEndpoints(fi *fileInfo, relPath string) []scanner.Endpoint {
	oaPath := expressPathToOpenAPI(filebasedRoutePath(relPath))
	var endpoints []scanner.Endpoint
	for _, exp := range findAllByType(fi.Root, "export_statement") {
		if ep, ok := buildFilebasedEndpoint(exp, fi, oaPath, relPath); ok {
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints
}
