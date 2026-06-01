//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 메서드의 @Path에서 경로를 추출한다(상수/결합/정규식 해석)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodPath(m *sitter.Node, fi *fileInfo) string {
	ann := findAnnotation(m, fi.src, AnnPath)
	if ann == nil {
		return ""
	}
	path := resolvePathArg(ann, fi.src, fi.imports, fi.absPath, fi.projectRoot)
	if path == "" {
		path = annotationElementValue(ann, fi.src, "value")
	}
	return path
}
