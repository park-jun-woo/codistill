//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 단일 URL 엔트리에서 View 유형별로 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildSingleURLEntryEndpoints builds endpoints for a single URL entry.
func buildSingleURLEntryEndpoints(entry urlEntry, viewsets []viewsetInfo, apiviews []apiviewInfo, funcViews []funcViewInfo, serializers map[string]serializerInfo, routerRegs map[string][]routerRegistration) []scanner.Endpoint {
	if entry.includeRouterVar != "" {
		return buildRouterRefEndpoints(entry, viewsets, serializers, routerRegs)
	}
	if len(entry.methodViews) > 0 {
		return buildRestPathEndpoints(entry)
	}
	viewName := resolveViewName(entry.viewName)
	if viewName == "" {
		return nil
	}
	if vs := findViewSet(viewsets, viewName); vs != nil {
		return buildURLMappedViewSetEndpoints(entry, vs, serializers)
	}
	if av := findAPIView(apiviews, viewName); av != nil {
		return buildAPIViewEndpoints(entry, av, serializers)
	}
	if fv := findFuncView(funcViews, viewName); fv != nil {
		return buildFuncViewEndpoints(entry, fv)
	}
	return buildPlainViewEndpoints(entry, viewName)
}
