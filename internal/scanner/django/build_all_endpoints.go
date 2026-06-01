//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 3-pass 파싱 결과를 결합하여 전체 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildAllEndpoints runs 3-pass extraction and combines endpoints.
func buildAllEndpoints(files []fileInfo) []scanner.Endpoint {
	urlEntries := collectURLs(files)
	routerRegs := extractRouterRegistrations(files)
	wrappers := collectRegisterWrappers(files)
	routerRegs = append(routerRegs, extractWrapperRegisterCalls(files, wrappers)...)
	classes := buildClassIndex(files)
	viewsets := collectViewSets(files, classes)
	apiviews := collectAPIViews(files, classes)
	funcViews := collectFuncViews(files)
	serializers := extractSerializers(files)

	regsByVar := routerRegsByVar(routerRegs)
	wired := wiredRouterKeys(urlEntries)

	var endpoints []scanner.Endpoint
	endpoints = append(endpoints, buildRouterEndpoints(unwiredRouterRegs(routerRegs, wired), viewsets, serializers)...)
	endpoints = append(endpoints, buildURLEntryEndpoints(urlEntries, viewsets, apiviews, funcViews, serializers, regsByVar)...)
	return endpoints
}
