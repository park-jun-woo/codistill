//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 등록의 receiver 변수명을 blueprint/namespace prefix로 해석한다
package flask

// resolveRegPrefix maps a resource registration's receiver variable (the base
// before .add_resource, or a configure_api blueprint variable) to a URL prefix.
// A blueprint prefix takes precedence; otherwise a flask_restx namespace prefix
// is used. An unknown receiver (e.g. a plain "api" Api object) resolves to no prefix.
func resolveRegPrefix(regVar string, bpPrefixes blueprintPrefix, nsPrefixes namespacePrefix) string {
	if p, ok := bpPrefixes[regVar]; ok {
		return p
	}
	if p, ok := nsPrefixes[regVar]; ok {
		return p
	}
	return ""
}
