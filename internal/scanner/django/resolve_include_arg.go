//ff:func feature=scan type=extract control=sequence topic=django
//ff:what include()의 인자를 router.urls 참조 또는 문자열 모듈명으로 해석한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// resolveIncludeArg interprets the argument list of an include(...) call. A
// `router.urls` attribute argument marks the entry as a router reference; an inline
// list argument (`include([path(...), *router.urls])`) is parsed into inline child
// entries; a string argument is taken as the included module path; and a bare
// identifier argument (`include(api_urls)`) is recorded as a file-scoped local list
// variable reference (includeLocalVar).
func resolveIncludeArg(entry *urlEntry, innerArgs *sitter.Node, src []byte) {
	if innerArgs == nil {
		return
	}
	if rv := routerVarFromURLsAttr(innerArgs, src); rv != "" {
		entry.includeRouterVar = rv
		return
	}
	if listNode := findChildByType(innerArgs, "list"); listNode != nil {
		entry.includeInline = parsePathCallsInList(listNode, src)
		return
	}
	if mod := firstStringArg(innerArgs, src); mod != "" {
		entry.includeModule = mod
		return
	}
	// A bare identifier argument (`include(api_urls)`) references a file-scoped local
	// list variable; record its name for later resolution against the file index.
	entry.includeLocalVar = firstIdentArg(innerArgs, src)
}
