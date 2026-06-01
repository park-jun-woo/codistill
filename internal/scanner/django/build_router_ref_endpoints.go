//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what router.urls 참조 entry의 누적 prefix를 각 register에 주입하여 CRUD를 전개한다
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildRouterRefEndpoints expands a urlEntry that wires a router's URLs into the
// urlconf (include(router.urls) or urlpatterns = router.urls). The entry's
// accumulated pattern is the include prefix; it is normalized (re_path regex
// named groups -> {name}) and composed with each registration's prefix, the
// registered ViewSet is expanded to CRUD endpoints, and the prefix's named-group
// path parameters (e.g. a nested router's "organizer"/"event") are inherited by
// every produced endpoint.
func buildRouterRefEndpoints(entry urlEntry, viewsets []viewsetInfo, serializers map[string]serializerInfo, routerRegs map[string][]routerRegistration) []scanner.Endpoint {
	normPrefix := djangoURLToOpenAPI(entry.pattern)
	prefixParams := extractURLParams(entry.pattern)
	var endpoints []scanner.Endpoint
	for _, reg := range routerRegs[entry.includeRouterVar] {
		vs := findViewSet(viewsets, resolveViewName(reg.viewsetName))
		if vs == nil {
			continue
		}
		prefixed := reg
		prefixed.prefix = combinePath(normPrefix, reg.prefix)
		eps := buildViewSetEndpoints(prefixed, vs, serializers)
		for i := range eps {
			prependPrefixParams(&eps[i], prefixParams)
		}
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
