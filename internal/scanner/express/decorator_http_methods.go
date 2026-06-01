//ff:type feature=scan type=model topic=express
//ff:what 데코레이터 라우팅용 HTTP 메서드 데코레이터명 → 메서드 매핑 및 컨트롤러 데코레이터명 집합
package express

// decoratorHTTPMethods maps method-decorator names (e.g. @Get, @Post) used by
// @n8n/decorators-style routing to their uppercase HTTP method.
var decoratorHTTPMethods = map[string]string{
	"Get":    "GET",
	"Post":   "POST",
	"Put":    "PUT",
	"Patch":  "PATCH",
	"Delete": "DELETE",
}

// controllerDecorators are class-level decorators that mark a routing
// controller in @n8n/decorators-style routing.
var controllerDecorators = map[string]bool{
	"RestController": true,
	"Controller":     true,
}
