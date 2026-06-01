//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_Success 테스트
package fastify

import "testing"

func TestExtractObjectRoute_Success(t *testing.T) {
	fi, calls := routeCalls(t, `server.route({ method: "GET", url: "/u", handler: h });`+"\n")
	inst := map[string]bool{"server": true}
	var got []routeInfo
	for _, c := range calls {
		if rs := extractObjectRoute(c, fi.Src, inst); rs != nil {
			got = rs
		}
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 route, got %d", len(got))
	}
	if got[0].Method != "GET" || got[0].Path != "/u" || got[0].Handler != "h" {
		t.Fatalf("route = %+v", got[0])
	}
}
