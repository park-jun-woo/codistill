//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_UrlNotString 테스트
package fastify

import "testing"

func TestExtractObjectRoute_UrlNotString(t *testing.T) {
	fi, calls := routeCalls(t, `server.route({ method: "GET", url: 123 });`+"\n")
	inst := map[string]bool{"server": true}
	for _, c := range calls {
		if rs := extractObjectRoute(c, fi.Src, inst); rs != nil {
			t.Fatalf("expected nil for non-string url, got %+v", rs)
		}
	}
}
