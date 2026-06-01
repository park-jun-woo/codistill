//ff:func feature=scan type=test control=sequence topic=express
//ff:what 테스트: 인자 없는 메서드 데코레이터(@Get)는 컨트롤러 prefix만으로 path를 합성한다
package express

import "testing"

func TestScan_DecoratorPrefixOnly(t *testing.T) {
	dir := t.TempDir()

	controller := `
import { Get, RestController } from '@n8n/decorators';

@RestController('/health')
export class HealthController {
	@Get()
	async check(req: any) {
		return;
	}
}
`
	writeFile(t, dir, "src/health.controller.ts", controller)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d: %v", len(result.Endpoints), result.Endpoints)
	}
	ep := result.Endpoints[0]
	if ep.Method != "GET" || ep.Path != "/health" {
		t.Errorf("expected GET /health, got %s %s", ep.Method, ep.Path)
	}
}
