//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 한 번의 fold 패스: 각 간선의 부모 prefix를 자식 effective prefix에 누적한다
package flask

// foldPass applies one accumulation pass over the register edges, prefixing each
// child's base contribution with its parent's current effective prefix. It
// reports whether any effective prefix changed, so the caller can iterate to a
// fixpoint. Edges with no resolvable parent (app/api receiver) are skipped here;
// their base contribution already sits in eff.
func foldPass(eff, base map[string]string, edges []registerEdge) bool {
	changed := false
	for _, e := range edges {
		if e.parentBP == "" {
			continue
		}
		parentPrefix, ok := eff[e.parentBP]
		if !ok {
			continue
		}
		want := combinePath(parentPrefix, base[e.childBP])
		if eff[e.childBP] != want {
			eff[e.childBP] = want
			changed = true
		}
	}
	return changed
}
