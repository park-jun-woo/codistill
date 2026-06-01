//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 뷰 래퍼 데코레이터의 첫 위치 인자를 실제 뷰로 언랩한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// tryUnwrapViewWrapper resolves the first positional argument of a view-wrapping
// decorator call (e.g. staff_member_required(view)) into entry. It reports
// whether the inner argument resolved to a plausible view (a name or an
// include); only then is entry replaced. A non-view first argument (e.g. a
// permission string or option list) leaves entry untouched and returns false so
// the caller can fall back to legacy behavior and keep the wrapper name.
func tryUnwrapViewWrapper(entry *urlEntry, arg *sitter.Node, src []byte) bool {
	inner := firstPositionalArgNode(arg)
	if inner == nil {
		return false
	}
	var inEntry urlEntry
	resolveSecondArg(&inEntry, inner, src)
	if inEntry.viewName == "" && !inEntry.isInclude {
		return false
	}
	inEntry.pattern = entry.pattern
	*entry = inEntry
	return true
}
