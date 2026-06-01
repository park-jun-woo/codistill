//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 같은 파일 내 include_in_schema=False 라우터 변수 집합을 구성한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveHiddenRouters builds the set of router variable names that are hidden
// from the OpenAPI schema, i.e. declared with APIRouter(include_in_schema=False)
// or included via include_router(child, include_in_schema=False). Routes attached
// to a hidden router variable must be skipped during collection.
func resolveHiddenRouters(root *sitter.Node, src []byte, routerSubclassNames map[string]bool) map[string]bool {
	hidden := make(map[string]bool)
	for _, r := range findRouterAssignments(root, src, routerSubclassNames) {
		if r.hidden {
			hidden[r.varName] = true
		}
	}
	for _, inc := range findIncludeRouterCalls(root, src) {
		if inc.hidden {
			hidden[inc.childVar] = true
		}
	}
	return hidden
}
