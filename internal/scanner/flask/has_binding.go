//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what from-import 바인딩 맵에 특정 로컬명이 있는지 확인한다
package flask

// hasBinding reports whether a local name was bound by a from-import in the file,
// distinguishing an imported blueprint receiver from a plain app/api object.
func hasBinding(bindings map[string]fromImportBinding, name string) bool {
	_, ok := bindings[name]
	return ok
}
