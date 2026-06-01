//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestContainsPrefixCall 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestContainsPrefixCall(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "hit.ts", "app.setGlobalPrefix('api');")
	writeFile(t, dir, "ver.ts", "app.enableVersioning({});")
	writeFile(t, dir, "miss.ts", "export const x = 1;")
	if !containsPrefixCall(filepath.Join(dir, "hit.ts")) {
		t.Fatal("expected hit for setGlobalPrefix")
	}
	if !containsPrefixCall(filepath.Join(dir, "ver.ts")) {
		t.Fatal("expected hit for enableVersioning")
	}
	if containsPrefixCall(filepath.Join(dir, "miss.ts")) {
		t.Fatal("expected no hit for plain file")
	}
}
