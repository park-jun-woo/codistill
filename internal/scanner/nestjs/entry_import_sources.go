//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 엔트리 .ts 파일의 import 문에서 모듈 소스 문자열을 수집한다
package nestjs

import "os"

// entryImportSources parses the .ts file at path and returns the source module
// strings of its import statements (e.g. "@gauzy/core", "./app.module"). Used
// to discover workspace packages imported by the bootstrap entry. A missing or
// unparseable file yields nil. Unlike extractImports this keeps non-relative
// (package) sources, since workspace package names must be matched downstream.
func entryImportSources(path string) []string {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	astRoot, err := parseTypeScript(src)
	if err != nil {
		return nil
	}
	var sources []string
	for _, stmt := range findAllByType(astRoot, "import_statement") {
		strNode := findChildByType(stmt, "string")
		if strNode == nil {
			continue
		}
		sources = append(sources, unquoteTS(nodeText(strNode, src)))
	}
	return sources
}
