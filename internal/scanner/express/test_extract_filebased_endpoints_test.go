//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 파일기반 export const VERB 추출 테스트: GET/POST 추출, config 비추출(회귀 안전)
package express

import (
	"sort"
	"testing"
)

func TestExtractFilebasedEndpoints(t *testing.T) {
	src := []byte(`
import { foo } from "x"
export const GET = async (req, res) => { res.json({ ok: true }) }
export const POST = async (req, res) => { res.status(201).json({}) }
export const config = { routes: [] }
const helper = () => {}
`)
	fi := mustParse(t, src)
	eps := extractFilebasedEndpoints(fi, "src/api/store/products/route.ts")
	if len(eps) != 2 {
		t.Fatalf("got %d endpoints, want 2: %+v", len(eps), eps)
	}
	got := map[string]string{}
	for _, ep := range eps {
		got[ep.Method] = ep.Path
		if ep.Handler != ep.Method {
			t.Errorf("handler %q != method %q", ep.Handler, ep.Method)
		}
	}
	if got["GET"] != "/store/products" {
		t.Errorf("GET path = %q, want /store/products", got["GET"])
	}
	if got["POST"] != "/store/products" {
		t.Errorf("POST path = %q, want /store/products", got["POST"])
	}
	methods := []string{}
	for m := range got {
		methods = append(methods, m)
	}
	sort.Strings(methods)
	if methods[0] != "GET" || methods[1] != "POST" {
		t.Errorf("methods = %v, want [GET POST]", methods)
	}
}
