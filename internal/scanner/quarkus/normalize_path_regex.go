//ff:func feature=scan type=convert control=iteration dimension=1 topic=quarkus
//ff:what JAX-RS path 변수의 정규식(`{name:regex}`)을 `{name}`으로 축약한다
package quarkus

import "strings"

func normalizePathRegex(path string) string {
	var b strings.Builder
	for i := 0; i < len(path); {
		if path[i] != '{' {
			b.WriteByte(path[i])
			i++
			continue
		}
		end := strings.IndexByte(path[i:], '}')
		if end < 0 {
			b.WriteString(path[i:])
			break
		}
		end += i
		inner := path[i+1 : end]
		if colon := strings.IndexByte(inner, ':'); colon >= 0 {
			inner = inner[:colon]
		}
		b.WriteByte('{')
		b.WriteString(strings.TrimSpace(inner))
		b.WriteByte('}')
		i = end + 1
	}
	return b.String()
}
