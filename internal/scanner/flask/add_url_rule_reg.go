//ff:type feature=scan type=model topic=flask
//ff:what add_url_rule 등록 정보 구조체(path·핸들러·메서드·blueprint변수·앱루트강제 여부)
package flask

// addURLRuleReg holds a single route registration discovered from a
// X.add_url_rule(rule, endpoint, view[, methods=(...)]) call. Indico's
// IndicoBlueprint.add_url_rule registers a RequestHandler class as the view and
// declares HTTP methods via the methods= keyword. blueprintVar is the receiver
// variable used as a prefix key; appRoot is true when the rule started with the
// Indico "!" prefix, which forces the path to be resolved against the app root
// (the blueprint prefix is ignored).
type addURLRuleReg struct {
	rawPath      string   // URL rule string (raw Flask form, "!" already stripped)
	handler      string   // view identifier (RequestHandler class name or function)
	methods      []string // HTTP methods from methods= kwarg (empty -> default GET)
	blueprintVar string   // receiver variable name for prefix lookup ("" if none)
	appRoot      bool     // true when rule began with Indico "!" (skip blueprint prefix)
	file         string   // relative file path
	line         int      // line number (1-based)
}
