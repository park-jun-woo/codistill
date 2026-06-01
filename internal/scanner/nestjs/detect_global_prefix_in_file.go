//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 .ts 파일에서 setGlobalPrefix 호출을 찾아 (prefix, found)를 반환한다
package nestjs

import "os"

// detectGlobalPrefixInFile parses a single .ts file and looks for setGlobalPrefix.
// Returns (prefix, true) if found with a literal arg, ("", true) if found with
// a non-literal arg, or ("", false) if not found at all.
func detectGlobalPrefixInFile(path string) (string, bool) {
	src, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	astRoot, err := parseTypeScript(src)
	if err != nil {
		return "", false
	}
	calls := findAllByType(astRoot, "call_expression")
	found := false
	for _, call := range calls {
		if prefix, ok := trySetGlobalPrefix(call, astRoot, src); ok {
			return prefix, true
		}
		if hasSetGlobalPrefix(call, src) {
			found = true
		}
	}
	return "", found
}
