//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 파일 내에서 정의된 Blueprint 변수명 집합을 수집한다
package flask

// ownModuleBlueprints returns the set of blueprint variable names defined by
// Blueprint(...) constructors in this file, used to tell whether a
// register_blueprint receiver is itself a known blueprint (a nesting parent) or
// a plain app/api object.
func ownModuleBlueprints(fi fileInfo) map[string]struct{} {
	set := make(map[string]struct{})
	for _, bp := range collectBlueprints(fi.root, fi.src) {
		set[bp.varName] = struct{}{}
	}
	return set
}
