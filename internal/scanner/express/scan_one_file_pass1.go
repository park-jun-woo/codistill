//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 파일의 Pass 1: 파싱, 라우터 수집, import 해석, use-prefix 수집, Zod 스키마 수집
package express

import sitter "github.com/smacker/go-tree-sitter"

func scanOneFilePass1(path string, parsed map[string]*fileInfo, allRouters map[string]map[string]bool, absRoot string, aliases map[string]string, schemas map[string]*sitter.Node, schemaSrc map[string][]byte) []mountEntry {
	fi, err := parseFile(path)
	if err != nil {
		return nil
	}
	parsed[path] = fi
	expressRouterAliases := collectExpressRouterImports(fi)
	routers := collectRouters(fi, expressRouterAliases)
	allRouters[path] = routers
	// Zod 스키마 수집
	fileSchemas := collectZodSchemas(fi)
	for name, node := range fileSchemas {
		schemas[name] = node
		schemaSrc[name] = fi.Src
	}
	imports := resolveImports(fi, absRoot, aliases)
	mounts := resolveUsePrefixes(fi, routers, imports, absRoot, aliases)
	var entries []mountEntry
	for _, m := range mounts {
		entries = append(entries, mountEntry{
			prefix:       m.Prefix,
			varName:      m.VarName,
			filePath:     m.FilePath,
			sourceFile:   path,
			sourceRouter: m.SourceRouter,
		})
	}
	// 인라인 마운트(filePath="")와 크로스파일 마운트를 raw 그대로 반환한다.
	// prefix 합성은 전역 라우터 그래프(resolveRouterPrefixes)에서 (file,var) 단위로 수행한다.
	entries = append(entries, extractArrayRouteMounts(fi, routers, imports, path)...)
	return entries
}
