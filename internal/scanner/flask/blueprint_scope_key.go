//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 파일/모듈 스코프 + 변수명을 blueprint prefix 맵의 복합 키로 만든다
package flask

// blueprintScopeKey builds the composite key used to scope a blueprint prefix to
// a particular file (relPath) or module path. Same local variable name (e.g.
// "bp") defined in different files would otherwise collide in the flat name map;
// scoping by file/module disambiguates them. The NUL separator cannot appear in
// a file path or Python identifier, so the composite key never collides with a
// bare variable-name key.
func blueprintScopeKey(scope, varName string) string {
	return scope + "\x00" + varName
}
