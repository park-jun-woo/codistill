//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what Flask-AppBuilder API 클래스 본문에서 base_url(없으면 기본값)을 해석한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classBaseURL resolves the URL prefix of a Flask-AppBuilder API class. It reads
// the class's string attributes (classStringAttrs) and uses an explicit
// `base_url` first; otherwise it falls back to FAB's convention
// `/api/v1/<route_base or resource_name or lowercased class name>`. The
// class-name fallback strips a trailing "RestApi"/"Api" suffix and lowercases the
// remainder (FooApi -> foo, ChartRestApi -> chart) so paths match
// Flask-AppBuilder's auto-generated route_base.
func classBaseURL(classNode *sitter.Node, className string, src []byte) string {
	attrs := classStringAttrs(classNode, src)
	if base := attrs["base_url"]; base != "" {
		return base
	}
	seg := attrs["route_base"]
	if seg == "" {
		seg = attrs["resource_name"]
	}
	if seg == "" {
		seg = defaultRouteSegment(className)
	}
	return combinePath("/api/v1", seg)
}
