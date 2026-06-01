//ff:func feature=scan type=extract control=sequence topic=django
//ff:what call 타입 인자를 include() 또는 .as_view()로 해석한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveCallArg resolves a call-type second argument (include() or .as_view()).
func resolveCallArg(entry *urlEntry, arg *sitter.Node, src []byte) {
	callFunc := findChildByType(arg, "identifier")
	callAttr := findChildByType(arg, "attribute")

	if callFunc != nil && nodeText(callFunc, src) == "include" {
		entry.isInclude = true
		resolveIncludeArg(entry, findChildByType(arg, "argument_list"), src)
		return
	}

	if callAttr != nil {
		attrText := nodeText(callAttr, src)
		if strings.HasSuffix(attrText, ".as_view") {
			entry.viewName = strings.TrimSuffix(attrText, ".as_view")
			entry.methodActions = parseAsViewDict(arg, src)
			return
		}
		entry.viewName = attrText
		return
	}

	if callFunc == nil {
		return
	}
	name := nodeText(callFunc, src)
	// A known view-wrapping decorator (staff_member_required(view), ...) hides
	// the real view in its first positional argument. Unwrap it so the wrapper
	// name never leaks into the view name (and thus the operationId). Nested
	// wrappers and inner X.as_view(...) calls are handled by re-resolving the
	// inner argument as a second-position view argument.
	if isViewWrapper(name) && tryUnwrapViewWrapper(entry, arg, src) {
		return
	}
	entry.viewName = name
}
