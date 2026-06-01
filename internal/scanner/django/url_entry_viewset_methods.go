//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what URL의 as_view dict를 우선 사용하고 없으면 부모에서 메서드를 역산한다
package django

import "sort"

// urlEntryViewSetMethods resolves the HTTP methods for a path()-mapped ViewSet.
// When the URL declares an explicit as_view({"get": "list", ...}) dict it is the
// source of truth; otherwise it falls back to inferring methods from the
// ViewSet's parent mixins.
func urlEntryViewSetMethods(entry urlEntry, vs *viewsetInfo) []actionMethod {
	if len(entry.methodActions) == 0 {
		return viewsetCRUDMethods(vs)
	}
	httpMethods := make([]string, 0, len(entry.methodActions))
	for m := range entry.methodActions {
		httpMethods = append(httpMethods, m)
	}
	sort.Strings(httpMethods)

	var methods []actionMethod
	for _, m := range httpMethods {
		if httpMethod, ok := apiviewHTTPMethods[m]; ok {
			methods = append(methods, actionMethod{action: entry.methodActions[m], method: httpMethod})
		}
	}
	return methods
}
