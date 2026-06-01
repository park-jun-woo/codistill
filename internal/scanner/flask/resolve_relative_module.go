//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 상대 import(점 개수 + 하위경로)를 파일 패키지 기준 절대 모듈 경로로 해석한다
package flask

import "strings"

// resolveRelativeModule resolves a relative import (the dotted import_prefix plus
// an optional sub dotted_name) against the importing file's package. `dots` is the
// number of leading dots: one dot means "this package", N dots climb N-1 levels
// above it. The package is the directory module of the importing file (its own
// module path with the final component dropped for non-package files).
func resolveRelativeModule(filePkg string, dots int, sub string) string {
	parts := []string{}
	if filePkg != "" {
		parts = strings.Split(filePkg, ".")
	}
	up := dots - 1
	if up > len(parts) {
		up = len(parts)
	}
	base := parts[:len(parts)-up]
	if sub != "" {
		base = append(base, strings.Split(sub, ".")...)
	}
	return strings.Join(base, ".")
}
