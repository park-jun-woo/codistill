//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveConstStringIdentifier_Fallback 테스트 (비문자열 const→원시 식별자 폴백)
package nestjs

import "testing"

func TestResolveConstStringIdentifier_Fallback(t *testing.T) {
	src := []byte(`
const C = foo();

@Controller(C)
export class CtrlController {
  @Get()
  all() {}
}
`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "ctrl.controller.ts", "/abs/ctrl.controller.ts", map[string]string{}, root, "/tmp")
	if !ok {
		t.Fatal("expected ok")
	}
	// non-string initializer cannot be resolved → raw identifier kept
	if ci.prefix != "C" {
		t.Fatalf("prefix: want raw %q got %q", "C", ci.prefix)
	}
}
