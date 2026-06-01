//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 파일의 use 임포트를 단축명(또는 별칭)→FQCN 맵으로 수집한다
package laravel

import (
	"strings"
)

// importMap collects a file's `use` imports into a map from the locally-visible
// short name (the class's last segment, or its alias when `use X\Y as Z`) to
// the fully-qualified class name. This lets the resolver turn a route's short
// controller name (e.g. UserController) back into its FQCN
// (App\Http\Controllers\Api\UserController), which PSR-4 then maps to an exact
// file path — recovering the precision the full-parse linear scan used to give.
func importMap(fi *fileInfo) map[string]string {
	out := make(map[string]string)
	for _, use := range findAllByType(fi.root, "namespace_use_clause") {
		qn := findChildByType(use, "qualified_name")
		var fqcn string
		if qn != nil {
			fqcn = strings.TrimSpace(strings.ReplaceAll(nodeText(qn, fi.src), " ", ""))
		} else if n := findChildByType(use, "name"); n != nil {
			fqcn = nodeText(n, fi.src)
		}
		if fqcn == "" {
			continue
		}
		short := lastNamespaceSegment(fqcn)
		if alias := useAlias(use, fi.src); alias != "" {
			short = alias
		}
		if short != "" {
			out[short] = fqcn
		}
	}
	return out
}
