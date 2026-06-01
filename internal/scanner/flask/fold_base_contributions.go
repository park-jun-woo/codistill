//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 각 blueprint의 기여 prefix(자기 prefix, register override가 있으면 그것)를 산출한다
package flask

// foldBaseContributions computes each blueprint's base contribution prefix: its
// own constructor url_prefix, replaced by an explicit url_prefix override on its
// register_blueprint call when present. This base is what gets prefixed by any
// parent during the topological fold.
func foldBaseContributions(own map[string]string, edges []registerEdge) map[string]string {
	base := make(map[string]string, len(own))
	for k, v := range own {
		base[k] = v
	}
	for _, e := range edges {
		if e.override != "" {
			base[e.childBP] = e.override
		}
	}
	return base
}
