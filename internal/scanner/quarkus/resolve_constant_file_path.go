//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 상수 참조의 className으로부터 정의 파일 경로를 import/동일패키지/동일파일 순으로 해석한다
package quarkus

func resolveConstantFilePath(className string, imports map[string]string, referrerPath, projectRoot string) string {
	if className == "" {
		return referrerPath
	}
	if fqcn, ok := imports[className]; ok {
		if filePath := resolveImportPath(projectRoot, fqcn); filePath != "" {
			return filePath
		}
	}
	if filePath := resolveSamePackageClass(referrerPath, className); filePath != "" {
		return filePath
	}
	return resolveSameFileClass(referrerPath, className, projectRoot)
}
