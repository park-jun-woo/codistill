//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what router.route("/x").all(handler) 체인형 all → 5개 HTTP 메서드 엔드포인트 회귀 가드
package express

import "testing"

func TestRouterRouteAllChain(t *testing.T) {
	dir := t.TempDir()
	src := `
const express = require("express");
const router = express.Router();
router.route("/x").all(handler);
export default router;
`
	writeFile(t, dir, "router.ts", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 5 {
		t.Fatalf("expected 5 endpoints for router.route().all, got %d", len(result.Endpoints))
	}
	methods := map[string]bool{}
	for _, ep := range result.Endpoints {
		methods[ep.Method] = true
		if ep.Path != "/x" {
			t.Errorf("path: want /x, got %s", ep.Path)
		}
	}
	for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		if !methods[m] {
			t.Errorf("missing method %s", m)
		}
	}
}
