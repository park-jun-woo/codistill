//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what nested router include 접두사의 named-group path param을 엔드포인트 앞에 승계한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// prependPrefixParams inserts the include-prefix path parameters (e.g. the
// regex named groups "organizer"/"event" of a nested router include) ahead of an
// endpoint's own path params, so a nested ViewSet's "{pk}" detail route carries
// the inherited "organizer"/"event" parameters in declaration order. Params
// already present on the endpoint are not duplicated.
func prependPrefixParams(ep *scanner.Endpoint, params []urlParam) {
	if len(params) == 0 {
		return
	}
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
	existing := map[string]bool{}
	for _, p := range ep.Request.PathParams {
		existing[p.Name] = true
	}
	var head []scanner.Param
	for _, p := range params {
		if existing[p.name] {
			continue
		}
		oaType := djangoConverterToOpenAPI(p.converter)
		head = append(head, scanner.Param{Name: p.name, Type: oaType.Type})
	}
	ep.Request.PathParams = append(head, ep.Request.PathParams...)
}
