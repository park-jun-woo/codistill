//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스 노드에서 resourceInfo를 구성한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func buildResourceInfo(cls *sitter.Node, fi *fileInfo) resourceInfo {
	ri := resourceInfo{
		file:    fi.relPath,
		absFile: fi.absPath,
		imports: fi.imports,
	}
	nameNode := findChildByType(cls, "identifier")
	if nameNode != nil {
		ri.className = nodeText(nameNode, fi.src)
	}
	ri.prefix = extractClassPath(cls, fi)
	ri.roles = extractClassRoles(cls, fi.src)
	ri.endpoints = extractMethodEndpoints(cls, fi)
	return ri
}
