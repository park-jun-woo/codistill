//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what HTTP 메서드 문자열 슬라이스를 대문자로 정규화하고 순서 보존하며 중복을 제거한다
package django

import "strings"

// dedupUpperMethods uppercases each method string and returns a new slice with
// duplicates removed, preserving first-seen order.
func dedupUpperMethods(ms []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, m := range ms {
		u := strings.ToUpper(m)
		if !seen[u] {
			seen[u] = true
			out = append(out, u)
		}
	}
	return out
}
