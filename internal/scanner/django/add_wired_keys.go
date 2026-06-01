//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 한 모듈의 entry 목록에서 router.urls 참조 키를 집합에 추가한다
package django

// addWiredKeys records the module-scoped router key of every entry in a module that
// wires a router into the urlconf via router.urls. Inline include children
// (include([... *router.urls])) are scanned recursively so their splatted router
// refs are also marked wired and skipped by the flat (prefix-less) router pass.
func addWiredKeys(entries []urlEntry, keys map[string]bool) {
	for _, entry := range entries {
		if entry.includeRouterVar != "" {
			keys[entry.includeRouterVar] = true
		}
		if len(entry.includeInline) > 0 {
			addWiredKeys(entry.includeInline, keys)
		}
	}
}
