//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 register를 감싸는 wrapper 함수명을 수집한다
package django

// collectRegisterWrappersFromFile finds module-local helper functions that wrap
// a router register: a function with at least two parameters (prefix, viewset)
// whose body contains a `.register(...)` call. Their names are returned so that
// calls to them can be promoted to router registrations.
func collectRegisterWrappersFromFile(fi fileInfo) map[string]bool {
	wrappers := map[string]bool{}
	for _, funcDef := range findAllByType(fi.root, "function_definition") {
		nameNode := findChildByType(funcDef, "identifier")
		if nameNode == nil {
			continue
		}
		if funcParamCount(funcDef, fi.src) < 2 {
			continue
		}
		if !funcBodyCallsRegister(funcDef, fi.src) {
			continue
		}
		wrappers[nodeText(nameNode, fi.src)] = true
	}
	return wrappers
}
