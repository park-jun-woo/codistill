//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 중첩 register_blueprint 간선을 위상적으로 접어 자식 prefix에 부모 prefix를 누적한다
package flask

// foldBlueprintPrefixes accumulates nested blueprint prefixes. For a
// `parent.register_blueprint(child[, url_prefix=ov])` edge, the child's effective
// prefix becomes parentEffectivePrefix + (ov or child's own prefix). Edges whose
// parent is app/api (parentBP == "") only apply the override to the child. The
// fold iterates to a fixpoint so chains (grandparent->parent->child) propagate;
// a bounded pass count keeps cycles from looping forever (cyclic/unresolved
// parents conservatively keep their own prefix). own holds each blueprint's
// constructor url_prefix; the returned map holds the resolved effective prefix.
func foldBlueprintPrefixes(own map[string]string, edges []registerEdge) map[string]string {
	base := foldBaseContributions(own, edges)
	eff := make(map[string]string, len(base))
	for k, v := range base {
		eff[k] = v
	}
	for pass := 0; pass <= len(edges); pass++ {
		if !foldPass(eff, base, edges) {
			break
		}
	}
	return eff
}
