//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what router.register 목록을 모듈 스코프 라우터 변수 키로 그룹핑한다
package django

// routerRegsByVar groups router registrations by their module-scoped variable key
// (module + var). A urlEntry that wires `router.urls` into the urlconf looks up its
// registrations here so the include prefix can be injected per registration.
func routerRegsByVar(regs []routerRegistration) map[string][]routerRegistration {
	byVar := make(map[string][]routerRegistration)
	for _, reg := range regs {
		key := routerKey(reg.module, reg.routerVar)
		byVar[key] = append(byVar[key], reg)
	}
	return byVar
}
