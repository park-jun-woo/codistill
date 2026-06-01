//ff:type feature=scan type=model topic=flask
//ff:what 클래스 기반 라우트의 HTTP 메서드 def 정보 구조체
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classMethod holds a single HTTP method def found inside a class-based route.
type classMethod struct {
	name     string       // uppercase HTTP method (GET, POST, ...)
	line     int          // 1-based line of the def
	funcNode *sitter.Node // the function_definition node
}
