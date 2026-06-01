//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모듈별 URL 맵을 루트부터 전개하여 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildURLEntryEndpoints expands the per-module URL map (following include) and builds endpoints.
func buildURLEntryEndpoints(byModule map[string][]urlEntry, viewsets []viewsetInfo, apiviews []apiviewInfo, funcViews []funcViewInfo, serializers map[string]serializerInfo, routerRegs map[string][]routerRegistration) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, root := range findRootURLModules(byModule) {
		entries := expandURLModule(root, "", byModule, map[string]bool{})
		for _, entry := range entries {
			eps := buildSingleURLEntryEndpoints(entry, viewsets, apiviews, funcViews, serializers, routerRegs)
			endpoints = append(endpoints, eps...)
		}
	}
	return endpoints
}
