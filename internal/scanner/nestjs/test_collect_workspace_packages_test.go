//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectWorkspacePackages 테스트
package nestjs

import "testing"

func TestCollectWorkspacePackages(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "packages/core/package.json", `{"name":"@app/core"}`)
	writeFile(t, dir, "apps/api/package.json", `{"name":"@app/api"}`)
	m := collectWorkspacePackages(dir)
	if _, ok := m["@app/core"]; !ok {
		t.Fatalf("expected @app/core in map, got %v", m)
	}
	if _, ok := m["@app/api"]; !ok {
		t.Fatalf("expected @app/api in map, got %v", m)
	}
}
