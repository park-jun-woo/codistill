//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what main.ts→src/*.ts 순회로 enableVersioning(URI) 호출과 defaultVersion을 감지한다
package nestjs

// detectURIVersioning searches for app.enableVersioning({ type: VersioningType.URI })
// across the same candidate files as setGlobalPrefix (main.ts first, then src/**),
// so bootstrap calls located outside src/main.ts (e.g. Novu's src/bootstrap.ts)
// are found. It returns (true, defaultVersion) when URI versioning is enabled;
// defaultVersion is "" when absent or not a simple string literal.
func detectURIVersioning(root string) (bool, string) {
	candidates := collectPrefixCandidates(root)
	for _, path := range candidates {
		if enabled, defaultVersion := detectURIVersioningInFile(path); enabled {
			return true, defaultVersion
		}
	}
	return false, ""
}
