//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 .ts 파일에서 enableVersioning(URI) 호출을 찾아 (enabled, defaultVersion)을 반환한다
package nestjs

import "os"

// detectURIVersioningInFile parses a single .ts file and looks for
// app.enableVersioning({ type: VersioningType.URI }). Returns (true, defaultVersion)
// when found, or (false, "") otherwise.
func detectURIVersioningInFile(path string) (bool, string) {
	src, err := os.ReadFile(path)
	if err != nil {
		return false, ""
	}
	astRoot, err := parseTypeScript(src)
	if err != nil {
		return false, ""
	}
	calls := findAllByType(astRoot, "call_expression")
	for _, call := range calls {
		if enabled, defaultVersion := isEnableURIVersioning(call, src); enabled {
			return true, defaultVersion
		}
	}
	return false, ""
}
