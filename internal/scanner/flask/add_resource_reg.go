//ff:type feature=scan type=model topic=flask
//ff:what add_resource/configure_api 등록 정보 구조체(클래스명·path 목록·blueprint변수)
package flask

// addResourceReg holds a single Flask-RESTful resource registration discovered
// from an api.add_resource(Resource, path[, path2...]) call or a
// configure_api_from_blueprint(blueprint, [(path, Resource), ...]) tuple form.
// blueprintVar is the blueprint variable name used as a prefix key (empty for
// plain api.add_resource calls that carry no blueprint prefix).
type addResourceReg struct {
	className    string   // Resource subclass identifier
	paths        []string // one or more URL rule strings (raw Flask form)
	blueprintVar string   // blueprint variable name for prefix lookup ("" if none)
}
