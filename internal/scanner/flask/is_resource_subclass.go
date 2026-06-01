//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 상위클래스 목록이 Flask-RESTful Resource를 상속하는지 판정한다
package flask

import "strings"

// isResourceSubclass reports whether any superclass resolves to Flask-RESTful's
// Resource: the final dotted component is "Resource" (e.g. flask_restful.Resource),
// or an import alias maps the bare name back to "Resource".
func isResourceSubclass(supers []string, aliases importAlias) bool {
	for _, s := range supers {
		base := s
		if i := strings.LastIndex(s, "."); i >= 0 {
			base = s[i+1:]
		}
		if base == "Resource" {
			return true
		}
		if aliases[base] == "Resource" {
			return true
		}
	}
	return false
}
