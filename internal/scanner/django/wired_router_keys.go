//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urlconf에 router.urls로 연결된 라우터 키 집합을 수집한다
package django

// wiredRouterKeys returns the set of module-scoped router keys that are wired into
// the urlconf via `include(router.urls)` or `urlpatterns = router.urls`. These are
// expanded with their include prefix through the URL path, so the flat
// (prefix-less) router pass must skip them to avoid duplicate, mis-prefixed routes.
func wiredRouterKeys(byModule map[string][]urlEntry) map[string]bool {
	keys := make(map[string]bool)
	for _, entries := range byModule {
		addWiredKeys(entries, keys)
	}
	return keys
}
