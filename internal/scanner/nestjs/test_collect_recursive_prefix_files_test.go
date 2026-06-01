//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectRecursivePrefixFiles 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestCollectRecursivePrefixFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/lib/bootstrap/index.ts", "app.setGlobalPrefix('api');")
	writeFile(t, dir, "src/app.module.ts", "export const x = 1;")
	writeFile(t, dir, "src/node_modules/pkg/x.ts", "app.setGlobalPrefix('skip');")
	files := collectRecursivePrefixFiles(filepath.Join(dir, "src"))
	if len(files) != 1 {
		t.Fatalf("expected exactly 1 prefix file (node_modules skipped), got %v", files)
	}
	if filepath.Base(files[0]) != "index.ts" {
		t.Fatalf("expected nested index.ts, got %v", files)
	}
}
