//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectURIVersioning 테스트
package nestjs

import "testing"

func TestDetectURIVersioning(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", `app.enableVersioning({ type: VersioningType.URI });`)
	if enabled, _ := detectURIVersioning(dir); !enabled {
		t.Fatal("expected URI versioning detected")
	}
}
