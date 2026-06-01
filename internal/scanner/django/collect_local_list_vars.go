//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 파일 내 `name = [ ...path()... ]` 로컬 리스트 변수 대입을 변수명→[]urlEntry로 수집한다
package django

// collectLocalListVars indexes file-scoped local list-variable assignments of the
// form `name = [path(...), ...]` (excluding the `urlpatterns` assignment) into a
// map from variable name to its parsed urlEntry children. This index resolves
// include(varName) references that name a same-file local list variable, e.g.
// `api_urls = [...]; path("api/v1/", include(api_urls))` (see resolveLocalVarIncludes).
func collectLocalListVars(fi fileInfo) map[string][]urlEntry {
	index := make(map[string][]urlEntry)
	for _, node := range findAllByType(fi.root, "assignment") {
		leftNodes := childrenOfType(node, "identifier")
		if len(leftNodes) == 0 {
			continue
		}
		name := nodeText(leftNodes[0], fi.src)
		if name == "urlpatterns" {
			continue
		}
		listNode := findChildByType(node, "list")
		if listNode == nil {
			continue
		}
		entries := parsePathCallsInList(listNode, fi.src)
		if len(entries) == 0 {
			continue
		}
		index[name] = entries
	}
	return index
}
