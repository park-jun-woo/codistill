//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_ChainTypeBox 테스트 — withTypeProvider 체인 라우트 인식
package fastify

import "testing"

func TestExtractOneRoute_ChainTypeBox(t *testing.T) {
	fi, calls := routeCalls(t, `fastify.withTypeProvider<TypeBoxTypeProvider>().get("/segments", { schema }, handler);`+"\n")
	inst := map[string]bool{"fastify": true}
	var got *routeInfo
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			got = ri
		}
	}
	if got == nil {
		t.Fatal("expected chained route")
	}
	if got.Method != "GET" || got.Path != "/segments" {
		t.Fatalf("chain route = %+v", got)
	}
}
