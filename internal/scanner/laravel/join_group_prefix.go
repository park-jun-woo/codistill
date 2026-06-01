//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 바깥 prefix와 안쪽 prefix를 결합한다
package laravel

import "strings"

// joinGroupPrefix joins outer prefix with inner prefix.
func joinGroupPrefix(outer, inner string) string {
	outer = strings.Trim(outer, "/")
	inner = strings.Trim(inner, "/")
	if outer == "" {
		return inner
	}
	if inner == "" {
		return outer
	}
	return outer + "/" + inner
}
