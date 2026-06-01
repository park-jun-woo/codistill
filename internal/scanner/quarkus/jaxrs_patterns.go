//ff:type feature=scan type=model topic=quarkus
//ff:what JAX-RS 어노테이션 이름 상수
package quarkus

const (
	AnnPath         = "Path"
	AnnGET          = "GET"
	AnnPOST         = "POST"
	AnnPUT          = "PUT"
	AnnDELETE       = "DELETE"
	AnnPATCH        = "PATCH"
	AnnPathParam    = "PathParam"
	AnnQueryParam   = "QueryParam"
	AnnHeaderParam  = "HeaderParam"
	AnnFormParam    = "FormParam"
	AnnRestForm     = "RestForm"
	AnnBeanParam    = "BeanParam"
	AnnContext      = "Context"
	AnnDefaultValue = "DefaultValue"
	AnnConsumes     = "Consumes"
	AnnProduces     = "Produces"
	AnnRolesAllowed = "RolesAllowed"
	AnnAuthenticated = "Authenticated"
	AnnPermitAll    = "PermitAll"
	AnnValid        = "Valid"
	AnnNotNull      = "NotNull"
	AnnNotBlank     = "NotBlank"
	AnnNotEmpty     = "NotEmpty"
	AnnMin          = "Min"
	AnnMax          = "Max"
	AnnSize         = "Size"
	AnnEmail        = "Email"
	AnnJsonProperty = "JsonProperty"
)

var httpMethodAnnotations = map[string]string{
	AnnGET:    "GET",
	AnnPOST:   "POST",
	AnnPUT:    "PUT",
	AnnDELETE: "DELETE",
	AnnPATCH:  "PATCH",
}

var skipDirs = map[string]bool{
	"build":        true,
	"target":       true,
	".gradle":      true,
	".mvn":         true,
	".git":         true,
	"node_modules": true,
	".idea":        true,
	"out":          true,
}

var primitiveTypes = map[string]bool{
	"String": true, "int": true, "Integer": true,
	"long": true, "Long": true, "float": true, "Float": true,
	"double": true, "Double": true, "boolean": true, "Boolean": true,
	"byte": true, "Byte": true, "short": true, "Short": true,
	"char": true, "Character": true,
}

// injectedParamTypes는 JAX-RS/서블릿 프레임워크 주입 객체 타입 set이다.
// 이 타입들은 요청 본문이 아니라 컨테이너가 주입하는 객체이므로
// requestBody 후보에서 배제한다(어노테이션 누락 대비 안전망).
var injectedParamTypes = map[string]bool{
	"HttpServletRequest":      true,
	"HttpServletResponse":     true,
	"UriInfo":                 true,
	"SecurityContext":         true,
	"Request":                 true,
	"HttpHeaders":             true,
	"SseEventSink":            true,
	"Sse":                     true,
	"AsyncResponse":           true,
	"ContainerRequestContext": true,
	"ResourceContext":         true,
	"Configuration":           true,
	"Application":             true,
	"Providers":               true,
	"ServletContext":          true,
	"ServletConfig":           true,
}

var responseStatusMethods = map[string]string{
	"Response.ok":         "200",
	"Response.created":    "201",
	"Response.accepted":   "202",
	"Response.noContent":  "204",
	"Response.seeOther":   "303",
	"Response.notModified": "304",
	"Response.serverError": "500",
}
