//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what OS별 경로 구분자를 슬래시로 정규화한다
package flask

import "strings"

// filepathToSlash normalizes OS-specific separators to forward slashes so module
// derivation is platform-independent.
func filepathToSlash(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}
