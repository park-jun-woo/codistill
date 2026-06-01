//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what argument_list에서 router.urls 형태 attribute의 라우터 변수명을 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// routerVarFromURLsAttr returns the router variable name if the argument list's
// first attribute argument is a `<var>.urls` reference (e.g. include(router.urls)
// yields "router"). It returns "" when no such attribute is present.
func routerVarFromURLsAttr(args *sitter.Node, src []byte) string {
	for _, child := range childrenOfType(args, "attribute") {
		text := nodeText(child, src)
		if strings.HasSuffix(text, ".urls") {
			return strings.TrimSuffix(text, ".urls")
		}
	}
	return ""
}
