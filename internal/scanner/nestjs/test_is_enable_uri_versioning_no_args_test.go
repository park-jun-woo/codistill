//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestIsEnableURIVersioning_NoArgs 테스트
package nestjs

import "testing"

func TestIsEnableURIVersioning_NoArgs(t *testing.T) {
	src := []byte(`app.enableVersioning();`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if enabled, _ := isEnableURIVersioning(call, src); !enabled {
		t.Fatal("expected true for no-args default URI")
	}
}
