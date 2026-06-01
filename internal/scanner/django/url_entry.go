//ff:type feature=scan type=model topic=django
//ff:what URL 패턴 엔트리 구조체
package django

// urlEntry represents a single path() entry from urls.py.
type urlEntry struct {
	pattern       string // URL pattern e.g. "api/users/<int:pk>/"
	viewName      string // view reference e.g. "UserViewSet" or "views.health_check"
	isInclude     bool   // whether the second arg is include(...)
	includeModule string // module path for include("app.urls")
	// includeRouterVar marks an include/assignment that wires a router's URLs into
	// the urlconf (`include(router.urls)` or `urlpatterns = router.urls`). During URL
	// collection it is qualified to a module-scoped router key (see routerKey).
	includeRouterVar string
	// includeInline holds the parsed entries of an inline list passed to include(),
	// e.g. include([path(...), *router.urls]). When set (with isInclude=true), the
	// entry's pattern is the prefix and these children are expanded beneath it
	// instead of resolving a string module (see expandURLEntry).
	includeInline []urlEntry
	// includeLocalVar names a file-scoped local list variable passed to include(),
	// e.g. `api_urls = [...]; path("api/v1/", include(api_urls))`. The argument is a
	// bare identifier (not a string module nor router.urls), so it is resolved against
	// the file's local list-variable index during collection (see resolveLocalVarIncludes)
	// by replacing it with the variable's parsed entries as includeInline children.
	includeLocalVar string
	// methodActions maps an HTTP method (lowercase, e.g. "get") to its ViewSet
	// action (e.g. "list") as declared in as_view({"get": "list", ...}).
	methodActions map[string]string
	// methodViews maps an HTTP method (uppercase, e.g. "GET") to its view
	// reference, as declared in a method-keyword helper like
	// rest_path("messages", GET=get_messages, POST=send_message). The keyword
	// argument names are HTTP methods and the values are view functions.
	methodViews map[string]string
}
