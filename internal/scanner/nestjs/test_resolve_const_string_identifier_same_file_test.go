//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveConstStringIdentifier_SameFile 테스트 (@Controller(CONST)→값)
package nestjs

import "testing"

func TestResolveConstStringIdentifier_SameFile(t *testing.T) {
	src := []byte(`
const HEALTH_CHECK_ROUTE = 'health';
const PING_PATH = '/ping';

@Controller(HEALTH_CHECK_ROUTE)
export class HealthController {
  @Get(PING_PATH)
  ping() {}
}
`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "health.controller.ts", "/abs/health.controller.ts", map[string]string{}, root, "/tmp")
	if !ok {
		t.Fatal("expected ok")
	}
	if ci.prefix != "health" {
		t.Fatalf("prefix: want %q got %q", "health", ci.prefix)
	}
	if len(ci.endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(ci.endpoints))
	}
	if ci.endpoints[0].path != "/ping" {
		t.Fatalf("method path: want %q got %q", "/ping", ci.endpoints[0].path)
	}
}
