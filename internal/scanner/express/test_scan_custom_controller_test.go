//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 테스트: 커스텀 Controller 베이스 클래스의 this.route/this.get 라우트 추출 + 게이트 회귀 안전
package express

import "testing"

func TestScan_CustomController(t *testing.T) {
	dir := t.TempDir()

	controller := `
import { Controller } from '../controller';

export class TokenController extends Controller {
	constructor() {
		super();
		this.route({ method: 'get', path: '/tokens', handler: this.getTokens });
		this.get('/preview/html/:template', this.preview);
		this.post('/tokens', this.createToken);
	}
}

// extends 가 Controller 가 아닌 클래스는 추출 대상 아님(회귀 안전)
export class NotAController extends BaseThing {
	constructor() {
		this.get('/ignored', this.handler);
		this.route({ method: 'get', path: '/also-ignored' });
	}
}
`
	writeFile(t, dir, "src/lib/routes/token.ts", controller)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	for _, e := range []string{
		"GET /tokens",
		"GET /preview/html/{template}",
		"POST /tokens",
	} {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
	for _, e := range []string{"GET /ignored", "GET /also-ignored"} {
		if found[e] {
			t.Errorf("non-Controller class endpoint %s was extracted (gate broken)", e)
		}
	}
	if len(result.Endpoints) != 3 {
		t.Errorf("expected 3 endpoints, got %d: %v", len(result.Endpoints), found)
	}
}
