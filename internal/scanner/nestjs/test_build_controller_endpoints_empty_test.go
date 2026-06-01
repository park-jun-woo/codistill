//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildControllerEndpoints_Empty 테스트
package nestjs

import "testing"

func TestBuildControllerEndpoints_Empty(t *testing.T) {
	cwf := controllerWithFile{info: controllerInfo{prefix: "x"}}
	eps, reqs := buildControllerEndpoints("", false, "", cwf, "", 0)
	if len(eps) != 0 || len(reqs) != 0 {
		t.Fatal("expected empty")
	}
}
