//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 register wrapper 함수 호출을 routerRegistration으로 수집한다
package django

// extractWrapperRegisterCalls finds calls to known register-wrapper helpers
// across all files and returns them as router registrations. The wrappers set
// is computed beforehand (collectRegisterWrappers) so that calls in any file
// can be promoted.
func extractWrapperRegisterCalls(files []fileInfo, wrappers map[string]bool) []routerRegistration {
	if len(wrappers) == 0 {
		return nil
	}
	var regs []routerRegistration
	for _, fi := range files {
		regs = append(regs, extractWrapperRegisterCallsFromFile(fi, wrappers)...)
	}
	return regs
}
