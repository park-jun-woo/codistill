//ff:func feature=scan type=extract control=sequence topic=django
//ff:what ViewSet의 CRUD 메서드를 반환한다(전이 해소된 값 우선, 없으면 직접 부모 역산)
package django

// viewsetCRUDMethods returns a ViewSet's CRUD HTTP methods. It prefers the
// transitively resolved methods filled at collection time (vs.methods); when those
// are absent (e.g. a viewsetInfo constructed directly without a class index) it
// falls back to the direct-parent resolution.
func viewsetCRUDMethods(vs *viewsetInfo) []actionMethod {
	if len(vs.methods) > 0 {
		return vs.methods
	}
	return resolveViewSetMethods(vs.parents)
}
