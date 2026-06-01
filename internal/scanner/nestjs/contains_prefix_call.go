//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what .ts 파일 본문에 setGlobalPrefix/enableVersioning 호출 텍스트가 있는지 검사한다
package nestjs

import (
	"os"
	"strings"
)

// containsPrefixCall reports whether the file at path contains the text
// "setGlobalPrefix(" or "enableVersioning(". This is a cheap pre-filter used
// before AST parsing during recursive bootstrap-file search, so a monorepo's
// full src/** is not parsed wholesale. Unreadable files yield false.
func containsPrefixCall(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	s := string(data)
	return strings.Contains(s, "setGlobalPrefix(") || strings.Contains(s, "enableVersioning(")
}
