//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 모든 파일의 Blueprint 정의를 모듈식별(module+varName) → 정보로 수집한다
package flask

// collectBlueprintIdentities collects every Blueprint(...) definition across files
// keyed by its module-qualified identity (defining module + variable name). The
// module qualification keeps same-named "bp" blueprints in different packages
// distinct, which the flat variable-name map cannot do.
func collectBlueprintIdentities(files []fileInfo) map[string]blueprintInfo {
	out := make(map[string]blueprintInfo)
	for _, fi := range files {
		module := modulePathOf(fi.relPath)
		for _, bp := range collectBlueprints(fi.root, fi.src) {
			out[blueprintScopeKey(module, bp.varName)] = bp
		}
	}
	return out
}
