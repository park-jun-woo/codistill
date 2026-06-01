//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 파일 내 blueprint 변수 참조를 모듈식별(module+varName)로 해석한다
package flask

// resolveBlueprintIdentity maps a blueprint variable referenced in a file to the
// module-qualified identity of the blueprint it actually points at. If the file
// imported the name via `from <module> import <orig>`, the identity is that
// source module's blueprint; otherwise the blueprint is assumed defined in the
// referencing file's own module. The returned key matches the keys produced for
// blueprint definitions, so prefixes resolve to the right blueprint even when
// many packages reuse the local name "bp".
func resolveBlueprintIdentity(varName, fileModule string, bindings map[string]fromImportBinding) string {
	if b, ok := bindings[varName]; ok {
		return blueprintScopeKey(b.module, b.orig)
	}
	return blueprintScopeKey(fileModule, varName)
}
