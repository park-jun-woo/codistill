//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what 모듈 내 라우터 참조 entry의 변수명을 모듈 스코프 키로 인플레이스 보정한다
package django

// qualifyRouterVars rewrites each entry's bare router variable (e.g. "router") into
// a module-scoped router key so a `router.urls` reference matches register() calls
// in the same module. Entries without a router reference are left untouched. Inline
// include children (include([... *router.urls])) are qualified recursively so their
// splatted router refs resolve against the same module's registrations.
func qualifyRouterVars(entries []urlEntry, module string) {
	for i := range entries {
		if entries[i].includeRouterVar != "" {
			entries[i].includeRouterVar = routerKey(module, entries[i].includeRouterVar)
		}
		if len(entries[i].includeInline) > 0 {
			qualifyRouterVars(entries[i].includeInline, module)
		}
	}
}
