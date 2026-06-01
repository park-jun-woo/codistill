//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 전 파일을 1회 훑어 라우터 서브클래스를 수집하고, 하나라도 추가됐으면 true 반환
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectRouterSubclassesSweep runs a single fixed-point sweep over all parsed
// files, adding newly discovered router subclass names into subclasses. It
// reports whether any new name was added during the sweep. Entries with a nil
// root are skipped.
func collectRouterSubclassesSweep(roots []*sitter.Node, srcs [][]byte, subclasses map[string]bool) bool {
	changed := false
	for i := range roots {
		if roots[i] == nil {
			continue
		}
		if collectRouterSubclassesPass(roots[i], srcs[i], subclasses) {
			changed = true
		}
	}
	return changed
}
