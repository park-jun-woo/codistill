//ff:type feature=scan type=model topic=flask
//ff:what 중첩 register_blueprint 등록 간선(부모/자식 모듈식별 + override) 구조체
package flask

// registerEdge records a `parent.register_blueprint(child[, url_prefix=...])`
// nesting relation. parentBP and childBP are module-qualified blueprint
// identities (module + "\x00" + varName) so same-named "bp" across packages
// stay distinct; override is the url_prefix= on the register call (empty if none).
type registerEdge struct {
	parentBP string // module-qualified parent blueprint id, "" if receiver is app/api
	childBP  string // module-qualified child blueprint id
	override string // url_prefix override on the register call, "" if absent
}
