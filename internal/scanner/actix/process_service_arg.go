//ff:func feature=scan type=extract control=selection topic=actix
//ff:what web::scope/web::resource 또는 fn()->Scope 간접등록 인자를 분기 처리해 라우트를 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func processServiceArg(callExpr *sitter.Node, src []byte, prefix string, routes *[]builderRoute, handlerFuncs map[string]*handlerInfo, visited map[string]bool) {
	switch findCallRoot(callExpr, src) {
	case "web::scope":
		scopePrefix := extractScopePrefix(callExpr, src)
		collectServiceCalls(callExpr, src, joinPath(prefix, scopePrefix), routes, handlerFuncs, visited)
		collectRouteCalls(callExpr, src, joinPath(prefix, scopePrefix), routes)
	case "web::resource":
		resourcePath := extractResourcePath(callExpr, src)
		fullPath := joinPath(prefix, resourcePath)
		collectRouteCalls(callExpr, src, fullPath, routes)
		collectToCalls(callExpr, src, fullPath, routes)
	default:
		// Indirect registration: `.service(xxx_scope())` where xxx_scope is a
		// plain function returning a `web::scope(...).service(...)` chain. Look
		// the function up in the cross-file index and recurse into its returned
		// chain so its scope prefix is composed onto the current prefix.
		resolveIndirectServiceArg(callExpr, src, prefix, routes, handlerFuncs, visited)
	}
}
