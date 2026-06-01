//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 같은 파일 내부에 정의된 클래스를 해석한다
package quarkus

import "os"

func resolveSameFileClass(referrerPath, className, projectRoot string) string {
	src, err := os.ReadFile(referrerPath)
	if err != nil {
		return ""
	}
	root, err := parseJava(src)
	if err != nil {
		return ""
	}
	classes := findAllByType(root, "class_declaration")
	classes = append(classes, findAllByType(root, "interface_declaration")...)
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode != nil && nodeText(nameNode, src) == className {
			return referrerPath
		}
	}
	return ""
}
