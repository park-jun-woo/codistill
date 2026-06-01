//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 메서드 노드에서 HTTP 라우트를 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractOneRoute(m *sitter.Node, fi *fileInfo) (endpointInfo, bool) {
	method, found := extractHTTPMethod(m, fi.src)
	if !found {
		return endpointInfo{}, false
	}
	path := extractMethodPath(m, fi)
	nameNode := findChildByType(m, "identifier")
	handler := ""
	if nameNode != nil {
		handler = nodeText(nameNode, fi.src)
	}
	ep := endpointInfo{
		method:  method,
		path:    path,
		handler: handler,
		file:    fi.relPath,
		line:    int(m.StartPoint().Row) + 1,
	}
	ep.roles = extractRolesFromNode(m, fi.src)
	extractMethodParams(m, fi.src, &ep, fi.imports, fi.absPath, fi.projectRoot)
	extractReturnInfo(m, fi.src, &ep)
	extractResponseStatus(m, fi.src, &ep)
	return ep, true
}
