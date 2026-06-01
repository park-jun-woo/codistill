//ff:type feature=scan type=model topic=flask
//ff:what 클래스 기반 라우트 카탈로그 타입(클래스명 → 메서드/노드/파일)
package flask

import sitter "github.com/smacker/go-tree-sitter"

// resourceClassInfo holds a discovered class-based route definition.
// Class-based scanners (Flask-RESTful / flask_restx / Flask-AppBuilder) collect
// these into a catalog keyed by class name, then resolve registrations
// (api.add_resource / appbuilder.add_api) against it.
type resourceClassInfo struct {
	name    string        // class name
	file    string        // relative file path
	methods []classMethod // HTTP method defs found in the class body
	node    *sitter.Node  // the class_definition node
}
