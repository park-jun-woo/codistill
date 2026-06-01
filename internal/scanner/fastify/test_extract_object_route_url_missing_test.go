//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_UrlMissing 테스트
package fastify

import "testing"

func TestExtractObjectRoute_UrlMissing(t *testing.T) {
	fi, calls := routeCalls(t, `server.route({ method: "GET" });`+"\n")
	inst := map[string]bool{"server": true}
	for _, c := range calls {
		if rs := extractObjectRoute(c, fi.Src, inst); rs != nil {
			t.Fatalf("expected nil for missing url, got %+v", rs)
		}
	}
}
