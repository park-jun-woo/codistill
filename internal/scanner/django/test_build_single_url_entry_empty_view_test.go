//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildSingleURLEntry_EmptyView 테스트
package django

import "testing"

func TestBuildSingleURLEntry_EmptyView(t *testing.T) {
	eps := buildSingleURLEntryEndpoints(urlEntry{viewName: ""}, nil, nil, nil, nil, nil)
	if eps != nil {
		t.Fatalf("expected nil for empty view name, got %+v", eps)
	}
}
