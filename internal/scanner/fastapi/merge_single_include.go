//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 include_router 호출에서 import된 라우터의 prefix를 병합한다
package fastapi

// mergeSingleInclude processes a single include_router call. If the childVar
// is imported from another file, the child's prefix is looked up from the
// global map and merged with the parent and extra prefix.
func mergeSingleInclude(fi *fileInfo, inc includeCall, importMap map[string]string, globalPrefixes map[string]map[string]string) {
	childVar := inc.childVar

	if _, local := fi.prefixes[childVar]; local {
		return
	}

	lookupKey := childVar
	if inc.childModule != "" {
		lookupKey = inc.childModule + "." + inc.childVar
	}
	srcFile := importMap[lookupKey]
	if srcFile == "" {
		return
	}
	srcPrefixes := globalPrefixes[srcFile]
	if srcPrefixes == nil {
		return
	}

	childKey := resolveImportOrigName(fi.imports, childVar)
	childPrefix, ok := srcPrefixes[childKey]
	if !ok {
		childPrefix = srcPrefixes[childVar]
	}
	parentPrefix := fi.prefixes[inc.parentVar]
	extra := resolveIfVariable(fi.root, inc.extraPrefix, fi.src)
	merged := joinPath(parentPrefix, extra, childPrefix)

	fi.prefixes[childVar] = merged
	if globalPrefixes[fi.absPath] == nil {
		globalPrefixes[fi.absPath] = make(map[string]string)
	}
	globalPrefixes[fi.absPath][childVar] = merged
}
