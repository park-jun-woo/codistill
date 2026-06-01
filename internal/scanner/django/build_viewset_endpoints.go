//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Router 등록된 ViewSet에서 CRUD + 커스텀 액션 엔드포인트를 생성한다
package django

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildViewSetEndpoints builds endpoints for a router-registered ViewSet.
func buildViewSetEndpoints(reg routerRegistration, vs *viewsetInfo, serializers map[string]serializerInfo) []scanner.Endpoint {
	prefix := "/" + strings.TrimLeft(reg.prefix, "/")
	methods := viewsetCRUDMethods(vs)

	var endpoints []scanner.Endpoint
	for _, am := range methods {
		ep := buildViewSetMethodEndpoint(prefix, am, vs, serializers)
		endpoints = append(endpoints, ep)
	}
	for _, ai := range vs.actions {
		eps := buildViewSetActionEndpoints(prefix, ai, vs)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
