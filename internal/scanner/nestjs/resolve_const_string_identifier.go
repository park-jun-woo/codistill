//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 데코레이터 경로 인자가 가리키는 const 문자열 상수를 해석한다
package nestjs

import (
	"os"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveConstStringIdentifier resolves a decorator path argument that is a
// simple identifier referencing a `const X = '<string>'` declaration (e.g.
// @Controller(HEALTH_CHECK_ROUTE) with const HEALTH_CHECK_ROUTE = 'health').
// It first searches the same file, then follows the import path of the
// identifier. Returns ("", false) when arg is not a simple identifier or the
// const cannot be resolved to a string value (caller keeps the raw text).
//
// Enum member expressions ("RouteKey.Asset") are handled by resolveEnumPathArg;
// this resolver is the complementary branch for dot-free identifiers.
func resolveConstStringIdentifier(arg string, root *sitter.Node, src []byte,
	filePath string, imports map[string]string, projectRoot string) (string, bool) {
	if !isSimpleConstIdent(arg) {
		return "", false
	}
	// 1. same file
	if v, ok := lookupConstString(root, src, arg); ok {
		return v, true
	}
	// 2. follow import path of the identifier
	importPath, ok := imports[arg]
	if !ok {
		return "", false
	}
	absPath := resolveImportPath(filepath.Dir(filePath), importPath, projectRoot)
	if absPath == "" {
		return "", false
	}
	constSrc, err := os.ReadFile(absPath)
	if err != nil {
		return "", false
	}
	constRoot, err := parseTypeScript(constSrc)
	if err != nil {
		return "", false
	}
	return lookupConstString(constRoot, constSrc, arg)
}
