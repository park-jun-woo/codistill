//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what 모듈 경로에 서브모듈 세그먼트를 import 점 규칙에 맞게 이어붙인다
package fastapi

import "strings"

// joinModulePath appends a sub-module segment to a Python import module path,
// honoring relative-import dot semantics. A purely relative module ("." or "..")
// already ends in its package boundary, so the segment is concatenated directly
// (`.` + `app` => `.app`). Any other module gets a "." separator
// (`.pkg` + `app` => `.pkg.app`, `pkg` + `app` => `pkg.app`).
func joinModulePath(module, segment string) string {
	if strings.HasSuffix(module, ".") {
		return module + segment
	}
	return module + "." + segment
}
