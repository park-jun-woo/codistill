//ff:func feature=scan type=convert control=sequence topic=django
//ff:what 라우터 변수 참조를 모듈 스코프 키로 만든다(모듈\x00변수)
package django

// routerKey builds a module-scoped key for a router variable. Router variables
// (e.g. "router") are local to a urls.py module, so a register() call and the
// `router.urls` reference that wires it must be matched within the same module.
func routerKey(module, routerVar string) string {
	return module + "\x00" + routerVar
}
