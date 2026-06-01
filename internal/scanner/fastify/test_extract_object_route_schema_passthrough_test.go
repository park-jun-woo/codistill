//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractObjectRoute_SchemaPassthrough 테스트 (opts 객체가 extractJSONSchema로 전달됨)
package fastify

import "testing"

func TestExtractObjectRoute_SchemaPassthrough(t *testing.T) {
	src := `server.route({ method: "POST", url: "/u", schema: { body: { type: "object", properties: { x: { type: "string" } } } } });` + "\n"
	fi, calls := routeCalls(t, src)
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
	si := extractJSONSchema(got[0].Schema, fi.Src)
	if si == nil {
		t.Fatal("expected schema info from opts passthrough")
	}
	if si.Body == nil {
		t.Fatal("expected body section in schema")
	}
}
