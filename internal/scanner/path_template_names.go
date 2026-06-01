//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 경로 템플릿에서 {token} 이름들을 순서대로 추출한다
package scanner

import "strings"

func pathTemplateNames(path string) []string {
	var names []string
	for {
		open := strings.IndexByte(path, '{')
		if open < 0 {
			break
		}
		close := strings.IndexByte(path[open:], '}')
		if close < 0 {
			break
		}
		// Skip `${...}` interpolation leftovers (a `{` immediately preceded by
		// `$`): an unresolved template fragment like `${PREFIX}` must not be
		// promoted to a synthetic path parameter. See Phase145.
		if open > 0 && path[open-1] == '$' {
			path = path[open+close+1:]
			continue
		}
		name := path[open+1 : open+close]
		if name != "" {
			names = append(names, name)
		}
		path = path[open+close+1:]
	}
	return names
}
