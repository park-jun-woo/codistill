//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 엔트리가 import한 워크스페이스 패키지 src에서 prefix 설정 파일을 모은다
package nestjs

import (
	"path/filepath"
)

// resolveWorkspacePrefixFiles handles the cross-package case where the scan
// root is a monorepo root and the entry file (src/main.ts) only imports a
// bootstrap symbol from a workspace package (e.g. import { bootstrap } from
// '@gauzy/core') without calling setGlobalPrefix itself. It maps that package
// name to its directory via package.json#name (collectWorkspacePackages) and
// recurses the package's src/ for prefix-setting files. External npm packages
// (not present in the workspace map) are ignored, so only workspace-local
// packages are tracked. Returns nil when no entry, no imports, or no match.
func resolveWorkspacePrefixFiles(root string) []string {
	imports := entryImportSources(filepath.Join(root, "src", "main.ts"))
	if len(imports) == 0 {
		return nil
	}
	pkgDirs := collectWorkspacePackages(root)
	if len(pkgDirs) == 0 {
		return nil
	}
	var files []string
	for _, imp := range imports {
		dir, ok := pkgDirs[imp]
		if !ok {
			continue
		}
		files = append(files, collectRecursivePrefixFiles(filepath.Join(dir, "src"))...)
	}
	return files
}
