//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what FastAPI/APIRouter 인스턴스와 include_router 호출에서 접두사 체인을 해석한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveRouterPrefixes finds all router variable assignments and builds
// a map of variable name -> full prefix (including include_router chains).
// Variable references (e.g., prefix=API_STR) are resolved against same-file
// assignments.
func resolveRouterPrefixes(root *sitter.Node, src []byte, routerSubclassNames map[string]bool) map[string]string {
	routers := findRouterAssignments(root, src, routerSubclassNames)
	includes := findIncludeRouterCalls(root, src)

	prefixes := make(map[string]string)
	for _, r := range routers {
		prefixes[r.varName] = resolveIfVariable(root, r.prefix, src)
	}
	for _, inc := range includes {
		// childModule이 있으면(`pkg.router`) 자식은 다른 모듈에서 import된
		// 라우터다. 로컬 prefix 맵에 그 이름이 없으며, 같은 로컬명(`router`)을
		// 여러 서브모듈이 공유하면 prefixes[childVar]가 include마다 자기 자신을
		// 누적해 폭주한다. 크로스파일 합성은 merge/propagate 패스가 담당하므로
		// 여기서는 로컬 정의 include(childModule=="")만 해석한다.
		if inc.childModule != "" {
			continue
		}
		extra := resolveIfVariable(root, inc.extraPrefix, src)
		parentPrefix := prefixes[inc.parentVar]
		childPrefix := prefixes[inc.childVar]
		prefixes[inc.childVar] = joinPath(parentPrefix, extra, childPrefix)
	}
	return prefixes
}
