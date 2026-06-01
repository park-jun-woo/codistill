//ff:type feature=scan type=model topic=flask
//ff:what Flask-AppBuilder API 클래스 정보 구조체(클래스명/base_url/모델여부)
package flask

import sitter "github.com/smacker/go-tree-sitter"

// appbuilderAPIInfo holds a discovered Flask-AppBuilder API class
// (a BaseApi / ModelRestApi / *RestApi subclass). The Flask-AppBuilder route
// scanner composes each class's @expose paths (and, for ModelRestApi, the
// standard CRUD endpoints) against baseURL.
type appbuilderAPIInfo struct {
	name           string       // class name
	file           string       // relative file path
	baseURL        string       // resolved base_url prefix (e.g. /api/v1/foo)
	isModelRestApi bool         // true when the class derives from ModelRestApi/*RestApi
	node           *sitter.Node // the class_definition node
}
