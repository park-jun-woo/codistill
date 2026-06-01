//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 최상위 .route("<path>", ...) 호출이면 그 인자에서 App 직속 라우트를 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func collectTopLevelRouteCall(n *sitter.Node, fi *fileInfo, routes *[]builderRoute) {
	if !isTopLevelRouteCall(n, fi.src) {
		return
	}
	args := findChildByType(n, "arguments")
	if args == nil {
		return
	}
	appendRouteFromArgs(args, fi.src, "", routes)
}
