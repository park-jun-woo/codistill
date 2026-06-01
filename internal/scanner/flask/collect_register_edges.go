//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 모든 파일의 register_blueprint 호출을 모듈식별 중첩 간선으로 수집한다
package flask

// collectRegisterEdges scans every file for register_blueprint calls and turns
// each into a module-qualified nesting edge. Parent and child variable names are
// resolved against the registering file's own module and its from-import
// bindings, so the edge connects the actual blueprints involved even when many
// packages reuse the local name "bp". An app/api receiver (unknown blueprint)
// yields an empty parent, meaning only an override prefix applies to the child.
func collectRegisterEdges(files []fileInfo) []registerEdge {
	var edges []registerEdge
	for _, fi := range files {
		edges = append(edges, fileRegisterEdges(fi)...)
	}
	return edges
}
