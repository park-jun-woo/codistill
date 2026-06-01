//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what main.ts 우선, src/** 재귀 + 워크스페이스 패키지에서 prefix 설정 파일을 모은다
package nestjs

import (
	"path/filepath"
)

// collectPrefixCandidates returns .ts file paths that may contain a
// setGlobalPrefix/enableVersioning call, with src/main.ts first. It recurses
// src/** (and monorepo packages/*/src/**) using a cheap text pre-filter
// (containsPrefixCall) so only bootstrap files are AST-parsed, then appends
// cross-package candidates from workspace packages imported by the entry file
// (resolveWorkspacePrefixFiles), e.g. main.ts importing { bootstrap } from
// '@gauzy/core'. Duplicates are removed while preserving first-seen order.
func collectPrefixCandidates(root string) []string {
	mainPath := filepath.Join(root, "src", "main.ts")
	candidates := []string{mainPath}

	candidates = append(candidates, collectRecursivePrefixFiles(filepath.Join(root, "src"))...)
	candidates = append(candidates, collectMonorepoPrefixFiles(root)...)
	candidates = append(candidates, resolveWorkspacePrefixFiles(root)...)

	return dedupePaths(candidates)
}
