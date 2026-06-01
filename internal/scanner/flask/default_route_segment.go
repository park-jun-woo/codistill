//ff:func feature=scan type=convert control=selection topic=flask
//ff:what 클래스명에서 RestApi/Api 접미사를 제거해 기본 route 세그먼트를 만든다
package flask

import "strings"

// defaultRouteSegment lowercases a class name after stripping a trailing
// "RestApi" or "Api" suffix, matching Flask-AppBuilder's default route_base
// (FooApi -> foo, ChartRestApi -> chart). Used by classBaseURL when no explicit
// base_url / route_base / resource_name is declared.
func defaultRouteSegment(className string) string {
	name := className
	switch {
	case strings.HasSuffix(name, "RestApi"):
		name = name[:len(name)-len("RestApi")]
	case strings.HasSuffix(name, "Api"):
		name = name[:len(name)-len("Api")]
	}
	if name == "" {
		name = className
	}
	return strings.ToLower(name)
}
