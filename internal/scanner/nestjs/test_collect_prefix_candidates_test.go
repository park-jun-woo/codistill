//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestCollectPrefixCandidates 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestCollectPrefixCandidates(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", "x")
	// Nested bootstrap file with a prefix call must be reached by recursion.
	writeFile(t, dir, "src/lib/bootstrap/index.ts", "app.setGlobalPrefix('api');")
	cands := collectPrefixCandidates(dir)
	if len(cands) < 2 {
		t.Fatalf("expected main.ts + nested bootstrap, got %v", cands)
	}
	if filepath.Base(cands[0]) != "main.ts" {
		t.Fatalf("main.ts should be first: %v", cands)
	}
	found := false
	for _, c := range cands {
		if filepath.Base(c) == "index.ts" {
			found = true
		}
	}
	if !found {
		t.Fatalf("nested bootstrap index.ts should be collected: %v", cands)
	}
}
