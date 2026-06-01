//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 로컬 import 이름(별칭 포함)에 대응하는 원본 모듈 내 이름을 반환한다
package fastapi

// resolveImportOrigName returns the original (source-module) name for a local
// imported name. For `from m import router as foo`, resolveImportOrigName("foo")
// returns "router". When the name has no recorded alias, it returns the name
// unchanged so callers can fall back to identity lookup.
func resolveImportOrigName(imports []importInfo, name string) string {
	for _, imp := range imports {
		if imp.name == name && imp.origName != "" {
			return imp.origName
		}
	}
	return name
}
