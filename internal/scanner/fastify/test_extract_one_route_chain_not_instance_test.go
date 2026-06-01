//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_ChainNotInstance 테스트 — 비인스턴스 체인 리시버는 nil
package fastify

import "testing"

func TestExtractOneRoute_ChainNotInstance(t *testing.T) {
	fi, calls := routeCalls(t, `foo.bar().get("/x", { schema }, h);`+"\n")
	inst := map[string]bool{"fastify": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("non-instance chain should yield nil, got %+v", ri)
		}
	}
}
