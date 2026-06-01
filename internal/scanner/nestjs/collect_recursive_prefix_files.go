//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 디렉터리를 재귀 순회해 prefix 설정 호출이 있는 .ts 파일 경로를 모은다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// collectRecursivePrefixFiles walks dir recursively and returns .ts files whose
// text contains a setGlobalPrefix/enableVersioning call (containsPrefixCall).
// node_modules, dist and .git directories and .d.ts files are skipped. A
// missing dir yields nil. This reaches bootstrap files nested below src/ (e.g.
// src/lib/bootstrap/index.ts) that the former src/ direct-only scan missed.
func collectRecursivePrefixFiles(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			switch info.Name() {
			case "node_modules", "dist", ".git":
				return filepath.SkipDir
			}
			return nil
		}
		name := info.Name()
		if !strings.HasSuffix(name, ".ts") || strings.HasSuffix(name, ".d.ts") {
			return nil
		}
		if containsPrefixCall(path) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
