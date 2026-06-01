//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 한 모노레포 레이아웃 디렉터리의 각 패키지 src에서 prefix 설정 파일을 모은다
package nestjs

import (
	"os"
	"path/filepath"
)

// collectLayoutPrefixFiles scans a single monorepo layout dir (e.g. "packages"
// or "apps") under root and returns prefix-setting .ts files found by recursing
// each child package's src/ (collectRecursivePrefixFiles). An absent layout dir
// yields nil.
func collectLayoutPrefixFiles(root, layout string) []string {
	entries, err := os.ReadDir(filepath.Join(root, layout))
	if err != nil {
		return nil
	}
	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		pkgSrc := filepath.Join(root, layout, e.Name(), "src")
		files = append(files, collectRecursivePrefixFiles(pkgSrc)...)
	}
	return files
}
