//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what register_blueprint 호출에서 Blueprint prefix 전파를 파일 스코프 + 중첩 누적으로 해석한다
package flask

// resolveBlueprintPrefixes builds the blueprint prefix map consumed by route
// extraction. Blueprints are collected by module-qualified identity so the same
// local name (e.g. "bp") in different packages does not clobber one another;
// nested parent.register_blueprint(child) prefixes are folded topologically. The
// returned map carries two key forms: a bare variable-name key (last-write
// fallback for single-blueprint / hand-built cases and existing override tests)
// and file-scoped keys (relPath + var) so each route resolves to the blueprint
// it actually imports.
func resolveBlueprintPrefixes(files []fileInfo) blueprintPrefix {
	idInfo := collectBlueprintIdentities(files)
	own := make(map[string]string, len(idInfo))
	for id, info := range idInfo {
		own[id] = info.urlPrefix
	}
	eff := foldBlueprintPrefixes(own, collectRegisterEdges(files))

	prefixes := make(blueprintPrefix)
	for id, info := range idInfo {
		prefixes[info.varName] = eff[id]
	}
	for _, fi := range files {
		addFileScopedPrefixes(fi, eff, prefixes)
	}
	return prefixes
}
