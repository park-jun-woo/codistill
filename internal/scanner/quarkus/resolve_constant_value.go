//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 상수 참조(Class.FIELD 또는 동일파일 FIELD)를 static final 필드의 실제 값으로 해석한다
package quarkus

import (
	"os"
	"strings"
)

func resolveConstantValue(constRef string, imports map[string]string, referrerPath, projectRoot string) string {
	className := ""
	fieldName := constRef
	if idx := strings.LastIndex(constRef, "."); idx >= 0 {
		className = constRef[:idx]
		fieldName = constRef[idx+1:]
		if dot := strings.IndexByte(className, '.'); dot >= 0 {
			className = className[dot+1:]
		}
	}

	filePath := resolveConstantFilePath(className, imports, referrerPath, projectRoot)
	if filePath == "" {
		return ""
	}

	src, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	root, err := parseJava(src)
	if err != nil {
		return ""
	}

	fileImports := imports
	if filePath != referrerPath {
		fileImports = extractImports(root, src)
	}

	classes := findAllByType(root, "class_declaration")
	return lookupStaticFinalInClasses(classes, src, className, fieldName, fileImports, filePath, projectRoot)
}
