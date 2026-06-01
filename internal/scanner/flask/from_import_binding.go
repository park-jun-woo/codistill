//ff:type feature=scan type=model topic=flask
//ff:what from-import 바인딩(로컬명 → 원본모듈/원본명) 구조체
package flask

// fromImportBinding records that a local name in a file was bound by a
// `from <module> import <orig> [as <local>]` statement, so a blueprint variable
// referenced by route decorators can be resolved back to the blueprint defined
// in its source module (disambiguating same-named "bp" across packages).
type fromImportBinding struct {
	module string // resolved source module path, e.g. "controllers.trigger"
	orig   string // original name in the source module, e.g. "bp"
}
