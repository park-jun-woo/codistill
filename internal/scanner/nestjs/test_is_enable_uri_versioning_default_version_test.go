//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestIsEnableURIVersioning_DefaultVersion 테스트
package nestjs

import "testing"

func TestIsEnableURIVersioning_DefaultVersion(t *testing.T) {
	src := []byte(`app.enableVersioning({ type: VersioningType.URI, defaultVersion: '1' });`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	enabled, defaultVersion := isEnableURIVersioning(call, src)
	if !enabled {
		t.Fatal("expected URI versioning enabled")
	}
	if defaultVersion != "1" {
		t.Fatalf("expected defaultVersion '1', got %q", defaultVersion)
	}
}
