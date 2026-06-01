//ff:type feature=scan type=model topic=flask
//ff:what flask_restx Namespace 변수명 → prefix 매핑 타입
package flask

// namespacePrefix maps a flask_restx Namespace variable name to its resolved
// URL prefix. The prefix is established by api.add_namespace(ns, path="/x").
// Used by the @ns.route class-based route scanner to compose
// namespace_prefix + route_path.
type namespacePrefix map[string]string
