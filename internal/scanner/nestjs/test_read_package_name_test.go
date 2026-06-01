//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestReadPackageName 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestReadPackageName(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "package.json", `{"name":"@app/core","version":"1.0.0"}`)
	if name := readPackageName(filepath.Join(dir, "package.json")); name != "@app/core" {
		t.Fatalf("expected @app/core, got %q", name)
	}
	if name := readPackageName(filepath.Join(dir, "missing.json")); name != "" {
		t.Fatalf("expected empty for missing, got %q", name)
	}
}
