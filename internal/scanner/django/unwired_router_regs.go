//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what router.urls로 urlconf에 연결되지 않은 register만 남긴다(평탄 fallback 전개용)
package django

// unwiredRouterRegs returns the registrations whose router is NOT wired into the
// urlconf via router.urls. Wired registrations are expanded with their include
// prefix through the URL path; the remaining ones keep the legacy flat, prefix-less
// expansion as a fallback so routers with no discoverable urlconf still surface.
func unwiredRouterRegs(regs []routerRegistration, wired map[string]bool) []routerRegistration {
	var out []routerRegistration
	for _, reg := range regs {
		if wired[routerKey(reg.module, reg.routerVar)] {
			continue
		}
		out = append(out, reg)
	}
	return out
}
