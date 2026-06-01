//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_MethodArray 테스트
package fastify

import "testing"

func TestExtractObjectRoute_MethodArray(t *testing.T) {
	fi, calls := routeCalls(t, `server.route({ method: ["GET", "POST"], url: "/u" });`+"\n")
	inst := map[string]bool{"server": true}
	var got []routeInfo
	for _, c := range calls {
		if rs := extractObjectRoute(c, fi.Src, inst); rs != nil {
			got = rs
		}
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(got))
	}
	if got[0].Method != "GET" || got[1].Method != "POST" {
		t.Fatalf("methods = %s, %s", got[0].Method, got[1].Method)
	}
}
