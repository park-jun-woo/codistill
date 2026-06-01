//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestIsEnableURIVersioning_NotVersioning 테스트
package nestjs

import "testing"

func TestIsEnableURIVersioning_NotVersioning(t *testing.T) {
	src := []byte(`app.listen(3000);`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if enabled, _ := isEnableURIVersioning(call, src); enabled {
		t.Fatal("expected false")
	}
}
