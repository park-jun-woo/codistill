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
	classes = append(classes, findAllByType(root, "interface_declaration")...)
	val := lookupStaticFinalInClasses(classes, src, className, fieldName, fileImports, filePath, projectRoot)
	if val != "" {
		return val
	}
	// 바인딩 없는 단순 식별자(className=="")는 상속 체인(extends/implements)의 상수일 수 있다.
	if className == "" {
		return resolveInheritedConstant(classes, src, fieldName, fileImports, filePath, projectRoot, map[string]bool{filePath: true})
	}
	return ""
}
