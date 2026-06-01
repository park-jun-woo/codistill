//ff:type feature=scan type=model topic=flask
//ff:what 클래스명 → resourceClassInfo 카탈로그 맵 타입
package flask

// resourceClassCatalog maps a class name to its discovered class-based route info.
// Class-based scanners (Flask-RESTful / flask_restx / Flask-AppBuilder) populate
// this catalog, then resolve registrations (api.add_resource / appbuilder.add_api)
// against it.
type resourceClassCatalog map[string]resourceClassInfo
