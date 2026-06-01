//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 include_router 호출의 축적된 prefix를 원본 파일에 전파한다
package fastapi

// propagateSingleInclude propagates the accumulated prefix from an
// include_router call in fi to the source file where the router is defined.
// It returns true if the source file's prefix was changed.
func propagateSingleInclude(fi *fileInfo, inc includeCall, importMap map[string]string,
	fileByPath map[string]*fileInfo, origSnapshot map[string]map[string]string) bool {

	lookupKey := inc.childVar
	if inc.childModule != "" {
		lookupKey = inc.childModule + "." + inc.childVar
	}
	srcFile := importMap[lookupKey]
	if srcFile == "" {
		return false
	}
	srcFI := fileByPath[srcFile]
	if srcFI == nil {
		return false
	}
	origVar := findOriginalVarName(inc.childVar, fi.imports, srcFI)
	if origVar == "" {
		return false
	}

	// 이 include 호출 고유의 extra 기여분은 키워드 prefix 자체다.
	// 스냅샷 strip 재구성은 동일 파일 내 hop 순서에 따라 부모 누적이
	// 자식에 반영되지 않은 비일관 상태를 만들 수 있어(다단계 체인),
	// 직접 resolve한 extraPrefix를 쓴다.
	extraContrib := resolveIfVariable(fi.root, inc.extraPrefix, fi.src)

	// 갱신된 부모 prefix + extra + 소스 파일의 원래 prefix
	parentPrefix := fi.prefixes[inc.parentVar]
	srcOrigPrefix := ""
	if sp := origSnapshot[srcFile]; sp != nil {
		srcOrigPrefix = sp[origVar]
	}
	accumulated := joinPath(parentPrefix, extraContrib, srcOrigPrefix)

	if accumulated == "" || srcFI.prefixes[origVar] == accumulated {
		return false
	}
	srcFI.prefixes[origVar] = accumulated
	return true
}
