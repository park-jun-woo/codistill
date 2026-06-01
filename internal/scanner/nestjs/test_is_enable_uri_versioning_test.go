//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestIsEnableURIVersioning 테스트
package nestjs

import "testing"

func TestIsEnableURIVersioning(t *testing.T) {
	src := []byte(`app.enableVersioning({ type: VersioningType.URI });`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if enabled, _ := isEnableURIVersioning(call, src); !enabled {
		t.Fatal("expected true")
	}
}
