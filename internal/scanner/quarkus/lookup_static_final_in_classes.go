//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 클래스 노드 목록에서 이름이 일치하는 클래스의 static final 필드 값을 찾아 반환한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func lookupStaticFinalInClasses(classes []*sitter.Node, src []byte, className, fieldName string, fileImports map[string]string, filePath, projectRoot string) string {
	for _, cls := range classes {
		if !classMatchesName(cls, src, className) {
			continue
		}
		val := findStaticFinalField(cls, src, fieldName, fileImports, filePath, projectRoot)
		if val != "" {
			return val
		}
	}
	return ""
}
