//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_NotRouteProp 테스트 (route 아닌 멤버 호출은 무시)
package fastify

import "testing"

func TestExtractObjectRoute_NotRouteProp(t *testing.T) {
	fi, calls := routeCalls(t, `server.register({ method: "GET", url: "/u" });`+"\n")
	inst := map[string]bool{"server": true}
	for _, c := range calls {
		if rs := extractObjectRoute(c, fi.Src, inst); rs != nil {
			t.Fatalf("expected nil for non-route property, got %+v", rs)
		}
	}
}
