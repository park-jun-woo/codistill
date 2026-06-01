//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 클래스 인덱스를 따라 부모 체인을 전이적으로 walk하여 DRF 베이스의 CRUD 메서드를 모은다
package django

// resolveMethodsTransitive walks the inheritance graph of parents through idx and
// collects the HTTP methods of every DRF ViewSet base/mixin reached (e.g. a custom
// ModelCrudViewSet that extends ModelViewSet yields the full CRUD set). A visited
// set guards against cycles; with a nil idx this degrades to a direct-parent check,
// matching resolveViewSetMethods.
func resolveMethodsTransitive(parents []string, idx classIndex) []actionMethod {
	var methods []actionMethod
	seen := make(map[string]bool)
	visited := make(map[string]bool)
	queue := append([]string(nil), parents...)
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if ms, ok := viewsetMethods[name]; ok {
			methods = appendUnseenMethods(methods, ms, seen)
		}
		if visited[name] {
			continue
		}
		visited[name] = true
		queue = append(queue, idx[name]...)
	}
	return methods
}
