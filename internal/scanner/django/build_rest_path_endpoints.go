//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what rest_path(GET=v, POST=v2) 메서드-키워드 엔트리에서 메서드별 엔드포인트를 생성한다
package django

import (
	"sort"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildRestPathEndpoints builds one endpoint per HTTP method declared in a
// method-keyword routing helper (rest_path). Each keyword value is a view
// reference and the keyword name fixes the HTTP method, so view-internal
// dispatch need not be resolved.
func buildRestPathEndpoints(entry urlEntry) []scanner.Endpoint {
	openAPIPath := ensureLeadingSlash(djangoURLToOpenAPI(entry.pattern))
	urlParams := extractURLParams(entry.pattern)

	methods := make([]string, 0, len(entry.methodViews))
	for method := range entry.methodViews {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	var endpoints []scanner.Endpoint
	for _, method := range methods {
		ep := scanner.Endpoint{
			Method:  method,
			Path:    openAPIPath,
			Handler: resolveViewName(entry.methodViews[method]),
		}
		addPathParams(&ep, urlParams)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
