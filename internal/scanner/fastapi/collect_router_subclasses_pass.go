//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 한 파일의 class_definition들을 1회 훑어 새 라우터 서브클래스를 subclasses에 추가하고 변경여부 반환
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectRouterSubclassesPass scans the class definitions of a single file once,
// adding any newly discovered router subclass names into subclasses. It reports
// whether at least one new name was added during this pass.
func collectRouterSubclassesPass(root *sitter.Node, src []byte, subclasses map[string]bool) bool {
	changed := false
	for _, cls := range findAllByType(root, "class_definition") {
		name := routerSubclassName(cls, src, subclasses)
		if name == "" {
			continue
		}
		subclasses[name] = true
		changed = true
	}
	return changed
}
