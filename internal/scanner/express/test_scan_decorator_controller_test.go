//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 테스트: @RestController + @Get/@Post 데코레이터 라우팅 추출(@n8n/decorators 스타일)
package express

import "testing"

func TestScan_DecoratorController(t *testing.T) {
	dir := t.TempDir()

	controller := `
import { Get, Post, RestController } from '@n8n/decorators';

@RestController('/mfa')
export class MFAController {
	@Post('/enforce-mfa')
	async enforceMFA(req: any) {
		return;
	}

	@Post('/can-enable', {
		allowSkipMFA: true,
	})
	async canEnableMFA(req: any) {
		return;
	}

	@Get('/qr', {
		allowSkipMFA: true,
	})
	async getQRCode(req: any) {
		return;
	}
}
`
	writeFile(t, dir, "src/controllers/mfa.controller.ts", controller)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"POST /mfa/enforce-mfa",
		"POST /mfa/can-enable",
		"GET /mfa/qr",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
	if len(result.Endpoints) != 3 {
		t.Errorf("expected 3 endpoints, got %d: %v", len(result.Endpoints), found)
	}
}
