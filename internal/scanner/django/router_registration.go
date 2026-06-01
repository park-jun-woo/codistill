//ff:type feature=scan type=model topic=django
//ff:what Router 등록 정보 구조체
package django

// routerRegistration represents a router.register("prefix", ViewSet) call.
type routerRegistration struct {
	prefix       string
	viewsetName  string
	basename     string
	routerPrefix string // prefix from include() if any
	routerVar    string // the router variable name, e.g. "router" in router.register(...)
	module       string // dotted module path where the register() call lives
}
