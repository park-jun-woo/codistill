//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildControllerInfo_TemplatePath 템플릿리터럴 ${CONST} 보간 경로 해석 테스트 (Phase145)
package nestjs

import "testing"

func TestBuildControllerInfo_TemplatePath(t *testing.T) {
	src := []byte("\n" +
		"const PREFIX_APIV3_DATA = '/api/v3/data/:baseName';\n" +
		"\n" +
		"@Controller()\n" +
		"export class DataV3Controller {\n" +
		"  @Get(`${PREFIX_APIV3_DATA}/:modelId/records`)\n" +
		"  list() {}\n" +
		"}\n")
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	ci, ok := buildControllerInfo(cls, src, "data-v3.controller.ts", "/abs/data-v3.controller.ts", map[string]string{}, root, "/tmp")
	if !ok {
		t.Fatal("expected ok")
	}
	if len(ci.endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(ci.endpoints))
	}
	want := "/api/v3/data/:baseName/:modelId/records"
	if ci.endpoints[0].path != want {
		t.Fatalf("method path: want %q got %q", want, ci.endpoints[0].path)
	}
}
