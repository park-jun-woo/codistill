//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 register wrapper 함수명 집합을 수집한다
package django

// collectRegisterWrappers aggregates register-wrapper function names across all
// parsed files. Wrappers are identified module-locally but invoked anywhere, so
// the name set is shared globally.
func collectRegisterWrappers(files []fileInfo) map[string]bool {
	wrappers := map[string]bool{}
	for _, fi := range files {
		for name := range collectRegisterWrappersFromFile(fi) {
			wrappers[name] = true
		}
	}
	return wrappers
}
