//ff:func feature=scan type=extract control=iteration dimension=2 topic=quarkus
//ff:what 단순 식별자 상수를 상속 체인(extends/implements)의 정의 클래스/인터페이스에서 재귀 해석한다
package quarkus

import (
	"os"

	sitter "github.com/smacker/go-tree-sitter"
)

func resolveInheritedConstant(classes []*sitter.Node, src []byte, fieldName string, imports map[string]string, referrerPath, projectRoot string, visited map[string]bool) string {
	for _, cls := range classes {
		for _, superName := range collectSupertypeNames(cls, src) {
			superPath := resolveConstantFilePath(superName, imports, referrerPath, projectRoot)
			if superPath == "" || visited[superPath] {
				continue
			}
			visited[superPath] = true
			superSrc, err := os.ReadFile(superPath)
			if err != nil {
				continue
			}
			root, err := parseJava(superSrc)
			if err != nil {
				continue
			}
			superImports := extractImports(root, superSrc)
			superClasses := findAllByType(root, "class_declaration")
			superClasses = append(superClasses, findAllByType(root, "interface_declaration")...)
			if val := lookupStaticFinalInClasses(superClasses, superSrc, "", fieldName, superImports, superPath, projectRoot); val != "" {
				return val
			}
			if val := resolveInheritedConstant(superClasses, superSrc, fieldName, superImports, superPath, projectRoot, visited); val != "" {
				return val
			}
		}
	}
	return ""
}
