//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveWorkspacePrefixFiles_External 테스트
package nestjs

import "testing"

func TestResolveWorkspacePrefixFiles_External(t *testing.T) {
	dir := t.TempDir()
	// Imported package has no workspace package.json mapping -> external; must
	// not be tracked, so no candidates are returned.
	writeFile(t, dir, "src/main.ts", `import { bootstrap } from '@nestjs/core';
bootstrap();
`)
	if files := resolveWorkspacePrefixFiles(dir); len(files) != 0 {
		t.Fatalf("external package must not be tracked, got %v", files)
	}
}
