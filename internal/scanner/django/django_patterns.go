//ff:type feature=scan type=model topic=django
//ff:what Django + DRF 패턴 상수 및 ViewSet 매핑 타입 정의
package django

// skipDirs are directory names to skip when walking Python files.
var skipDirs = map[string]bool{
	"venv":          true,
	".venv":         true,
	"__pycache__":   true,
	".git":          true,
	"dist":          true,
	"node_modules":  true,
	".tox":          true,
	".mypy_cache":   true,
	".pytest_cache": true,
	"migrations":    true,
}

// actionMethod pairs a ViewSet action with its HTTP method and detail flag.
type actionMethod struct {
	action string
	method string
	detail bool
}

// viewsetMethods maps DRF ViewSet mixin parents to the HTTP methods they provide.
var viewsetMethods = map[string][]actionMethod{
	"ModelViewSet": {
		{action: "list", method: "GET", detail: false},
		{action: "create", method: "POST", detail: false},
		{action: "retrieve", method: "GET", detail: true},
		{action: "update", method: "PUT", detail: true},
		{action: "partial_update", method: "PATCH", detail: true},
		{action: "destroy", method: "DELETE", detail: true},
	},
	"ReadOnlyModelViewSet": {
		{action: "list", method: "GET", detail: false},
		{action: "retrieve", method: "GET", detail: true},
	},
	"CreateModelMixin": {
		{action: "create", method: "POST", detail: false},
	},
	"ListModelMixin": {
		{action: "list", method: "GET", detail: false},
	},
	"RetrieveModelMixin": {
		{action: "retrieve", method: "GET", detail: true},
	},
	"UpdateModelMixin": {
		{action: "update", method: "PUT", detail: true},
		{action: "partial_update", method: "PATCH", detail: true},
	},
	"DestroyModelMixin": {
		{action: "destroy", method: "DELETE", detail: true},
	},
}

// apiviewHTTPMethods are the HTTP method names that APIView subclasses can define.
var apiviewHTTPMethods = map[string]string{
	"get":     "GET",
	"post":    "POST",
	"put":     "PUT",
	"delete":  "DELETE",
	"patch":   "PATCH",
	"options": "OPTIONS",
	"head":    "HEAD",
}
