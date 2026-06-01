//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_AnonymousHandler 테스트 (arrow handler → (anonymous))
package fastify

import "testing"

func TestExtractObjectRoute_AnonymousHandler(t *testing.T) {
	fi, calls := routeCalls(t, `server.route({ method: "GET", url: "/u", handler: () => {} });`+"\n")
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
	if got[0].Handler != "(anonymous)" {
		t.Fatalf("handler = %q", got[0].Handler)
	}
}
