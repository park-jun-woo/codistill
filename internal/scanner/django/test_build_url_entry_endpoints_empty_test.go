//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildURLEntryEndpoints_Empty 테스트
package django

import "testing"

func TestBuildURLEntryEndpoints_Empty(t *testing.T) {
	eps := buildURLEntryEndpoints(map[string][]urlEntry{}, nil, nil, nil, nil, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
