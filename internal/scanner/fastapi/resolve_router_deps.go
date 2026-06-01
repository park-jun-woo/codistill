//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what APIRouter 생성자의 dependencies 인자에서 라우터 변수별 미들웨어 맵을 구성한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveRouterDeps builds a map of router variable name -> middleware
// (dependencies) extracted from APIRouter(dependencies=[...]) constructor calls.
func resolveRouterDeps(root *sitter.Node, src []byte, routerSubclassNames map[string]bool) map[string][]string {
	routers := findRouterAssignments(root, src, routerSubclassNames)
	deps := make(map[string][]string)
	for _, r := range routers {
		if len(r.middleware) > 0 {
			deps[r.varName] = r.middleware
		}
	}
	return deps
}
