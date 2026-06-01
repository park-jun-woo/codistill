//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestEntryImportSources 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestEntryImportSources(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", `import { bootstrap } from '@app/core';
import { AppModule } from './app.module';
bootstrap();
`)
	got := entryImportSources(filepath.Join(dir, "src", "main.ts"))
	has := func(s string) bool {
		for _, g := range got {
			if g == s {
				return true
			}
		}
		return false
	}
	if !has("@app/core") || !has("./app.module") {
		t.Fatalf("expected both import sources, got %v", got)
	}
}
