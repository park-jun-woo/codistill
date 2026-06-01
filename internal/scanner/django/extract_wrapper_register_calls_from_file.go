//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 register wrapper 함수 호출을 routerRegistration으로 수집한다
package django

// extractWrapperRegisterCallsFromFile finds calls to known register-wrapper
// helpers in a single file and returns them as router registrations.
func extractWrapperRegisterCallsFromFile(fi fileInfo, wrappers map[string]bool) []routerRegistration {
	var regs []routerRegistration
	for _, callNode := range findAllByType(fi.root, "call") {
		if reg := parseWrapperRegisterCall(callNode, fi, wrappers); reg != nil {
			regs = append(regs, *reg)
		}
	}
	return regs
}
