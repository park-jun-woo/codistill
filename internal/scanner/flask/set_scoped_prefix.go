//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what blueprint 식별자의 effective prefix가 있으면 파일 스코프 키로 등록한다
package flask

// setScopedPrefix writes the file-scoped prefix entry for a blueprint reference
// when the referenced blueprint identity has a resolved effective prefix. A
// missing identity (unknown import) is skipped so route extraction falls back to
// the bare variable-name key.
func setScopedPrefix(prefixes blueprintPrefix, relPath, varName string, eff map[string]string, id string) {
	if p, ok := eff[id]; ok {
		prefixes[blueprintScopeKey(relPath, varName)] = p
	}
}
