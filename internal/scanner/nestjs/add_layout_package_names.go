//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 한 레이아웃 디렉터리의 각 패키지 package.json#name을 결과 맵에 추가한다
package nestjs

import (
	"os"
	"path/filepath"
)

// addLayoutPackageNames scans a single monorepo layout dir (e.g. "packages")
// under root and records each child package's package.json#name -> directory
// into result (readPackageName). An absent layout dir or unnamed package is
// skipped.
func addLayoutPackageNames(root, layout string, result map[string]string) {
	entries, err := os.ReadDir(filepath.Join(root, layout))
	if err != nil {
		return
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		pkgDir := filepath.Join(root, layout, e.Name())
		if name := readPackageName(filepath.Join(pkgDir, "package.json")); name != "" {
			result[name] = pkgDir
		}
	}
}
