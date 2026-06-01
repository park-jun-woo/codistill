//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 한 파일이 참조 가능한 blueprint들(자기 정의 + from-import)을 파일 스코프 prefix 키로 등록한다
package flask

// addFileScopedPrefixes binds every blueprint a file can reference — its own
// Blueprint(...) definitions and its from-import bindings — to that blueprint's
// effective prefix under a file-scoped key (relPath + var). Route extraction
// consults this key first so the same local name "bp" in different packages
// resolves to the right prefix instead of the last-collected one.
func addFileScopedPrefixes(fi fileInfo, eff map[string]string, prefixes blueprintPrefix) {
	fileModule := modulePathOf(fi.relPath)
	for _, bp := range collectBlueprints(fi.root, fi.src) {
		setScopedPrefix(prefixes, fi.relPath, bp.varName, eff, blueprintScopeKey(fileModule, bp.varName))
	}
	for local, b := range collectFromImports(fi.root, fi.src, fi.relPath) {
		setScopedPrefix(prefixes, fi.relPath, local, eff, blueprintScopeKey(b.module, b.orig))
	}
}
