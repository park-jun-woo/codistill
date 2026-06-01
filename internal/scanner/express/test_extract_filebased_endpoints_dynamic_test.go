//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 동적 세그먼트 + export async function 형태 파일기반 추출 테스트: [id]/route.ts → /store/products/{id}
package express

import "testing"

func TestExtractFilebasedEndpointsDynamic(t *testing.T) {
	src := []byte(`
export async function GET(req, res) { res.json({}) }
export const DELETE = async (req, res) => { res.sendStatus(204) }
`)
	fi := mustParse(t, src)
	eps := extractFilebasedEndpoints(fi, "src/api/store/products/[id]/route.ts")
	if len(eps) != 2 {
		t.Fatalf("got %d endpoints, want 2: %+v", len(eps), eps)
	}
	for _, ep := range eps {
		if ep.Path != "/store/products/{id}" {
			t.Errorf("%s path = %q, want /store/products/{id}", ep.Method, ep.Path)
		}
	}
}
