//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 라우트의 라우터 변수에 대한 prefix를 파일 스코프 우선으로 조회한다
package flask

// lookupBlueprintPrefix resolves the URL prefix for a router variable referenced
// by a route in file relPath. A file-scoped key (relPath, varName) wins so the
// same local name (e.g. "bp") defined in different blueprint packages does not
// clobber one another in the flat name map; it falls back to the bare
// variable-name key for the common single-blueprint / hand-built-map cases.
func lookupBlueprintPrefix(bpPrefixes blueprintPrefix, relPath, varName string) string {
	if p, ok := bpPrefixes[blueprintScopeKey(relPath, varName)]; ok {
		return p
	}
	return bpPrefixes[varName]
}
