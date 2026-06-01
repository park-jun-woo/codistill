//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 단일 파일의 register_blueprint 호출을 모듈식별 중첩 간선으로 변환한다
package flask

// fileRegisterEdges resolves every register_blueprint call in one file into a
// module-qualified nesting edge. Child/parent variable names are resolved via the
// file's own blueprint definitions and from-import bindings; a receiver that is
// neither a known blueprint nor an imported one (a plain app/api object) yields
// an empty parent so only an override prefix applies.
func fileRegisterEdges(fi fileInfo) []registerEdge {
	fileModule := modulePathOf(fi.relPath)
	bindings := collectFromImports(fi.root, fi.src, fi.relPath)
	ownBPs := ownModuleBlueprints(fi)
	var edges []registerEdge
	for _, call := range findAllByType(fi.root, "call") {
		parentVar, childVar, override := parseRegisterBlueprintEdge(call, fi.src)
		if childVar == "" {
			continue
		}
		childID := resolveBlueprintIdentity(childVar, fileModule, bindings)
		parentID := ""
		if _, isBP := ownBPs[parentVar]; isBP || hasBinding(bindings, parentVar) {
			parentID = resolveBlueprintIdentity(parentVar, fileModule, bindings)
		}
		edges = append(edges, registerEdge{parentBP: parentID, childBP: childID, override: override})
	}
	return edges
}
