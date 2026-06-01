//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 리스트 내 list_splat(*x) 노드를 파싱하여 router.urls splat이면 라우터 참조 entry를 반환한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// parseListSplat interprets a `list_splat` element of a urlpatterns list (e.g.
// `*router.urls` or `*intake_urls`). When the splatted operand is a `<var>.urls`
// attribute it yields a router reference entry so the registered ViewSets are
// expanded as CRUD (wired by Phase158). A bare-identifier splat (`*intake_urls`)
// is left to module-level expansion and returns nil.
func parseListSplat(splatNode *sitter.Node, src []byte) *urlEntry {
	attr := findChildByType(splatNode, "attribute")
	if attr == nil {
		return nil
	}
	text := nodeText(attr, src)
	if !strings.HasSuffix(text, ".urls") {
		return nil
	}
	return &urlEntry{isInclude: true, includeRouterVar: strings.TrimSuffix(text, ".urls")}
}
