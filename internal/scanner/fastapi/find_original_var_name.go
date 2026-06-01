//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what import된 이름(별칭 포함)에 대응하는 원본 파일의 라우터 변수명을 찾는다
package fastapi

// findOriginalVarName finds the original variable name in the source file for
// an imported name. It first resolves any import alias (e.g. `import router as
// foo`) to the original source-module name via refImports, then confirms that
// name is a router variable defined in srcFI. Falls back to the imported name
// when there is no alias. Returns "" if neither matches a router variable.
func findOriginalVarName(importedName string, refImports []importInfo, srcFI *fileInfo) string {
	orig := resolveImportOrigName(refImports, importedName)
	if _, ok := srcFI.prefixes[orig]; ok {
		return orig
	}
	if _, ok := srcFI.prefixes[importedName]; ok {
		return importedName
	}
	return ""
}
