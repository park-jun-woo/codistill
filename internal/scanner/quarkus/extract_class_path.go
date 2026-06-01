//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스의 @Path에서 prefix 경로를 추출한다(상수/결합/정규식 해석)
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractClassPath(cls *sitter.Node, fi *fileInfo) string {
	ann := findAnnotation(cls, fi.src, AnnPath)
	if ann == nil {
		return ""
	}
	path := resolvePathArg(ann, fi.src, fi.imports, fi.absPath, fi.projectRoot)
	if path == "" {
		path = annotationElementValue(ann, fi.src, "value")
	}
	return path
}
