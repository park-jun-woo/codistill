//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildAllEndpoints_Empty 테스트
package nestjs

import "testing"

func TestBuildAllEndpoints_Empty(t *testing.T) {
	eps, reqs := buildAllEndpoints("", false, "", nil, "")
	if len(eps) != 0 || len(reqs) != 0 {
		t.Fatal("expected empty results")
	}
}
