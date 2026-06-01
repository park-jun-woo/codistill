//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestBuildControllerEndpoints_ArrayFanout 테스트 (배열 경로 → endpoint 복제)
package nestjs

import "testing"

func TestBuildControllerEndpoints_ArrayFanout(t *testing.T) {
	cwf := controllerWithFile{
		info: controllerInfo{
			prefix: "",
			endpoints: []endpointInfo{
				{method: "GET", handler: "list", paths: []string{"/a", "/b"}},
				{method: "POST", handler: "create", path: "/c"},
			},
		},
		absFile: "/src/x.controller.ts",
	}
	eps, _ := buildControllerEndpoints("", false, "", cwf, "", 0)
	if len(eps) != 3 {
		t.Fatalf("expected 3 endpoints after fan-out, got %d", len(eps))
	}
	got := map[string]string{}
	for _, e := range eps {
		got[e.Path] = e.Method
	}
	if got["/a"] != "GET" || got["/b"] != "GET" {
		t.Fatalf("array paths not fanned out: %+v", got)
	}
	if got["/c"] != "POST" {
		t.Fatalf("single path lost: %+v", got)
	}
}
