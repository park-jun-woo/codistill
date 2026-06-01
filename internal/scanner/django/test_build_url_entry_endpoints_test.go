//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildURLEntryEndpoints 테스트
package django

import "testing"

func TestBuildURLEntryEndpoints(t *testing.T) {
	byModule := map[string][]urlEntry{
		"app.urls": {
			{pattern: "health/", viewName: "health"},
		},
	}
	funcViews := []funcViewInfo{{name: "health", methods: []string{"GET"}, file: "v.py"}}
	eps := buildURLEntryEndpoints(byModule, nil, nil, funcViews, map[string]serializerInfo{}, nil)
	if len(eps) == 0 {
		t.Fatal("expected endpoints from URL entries")
	}
	if eps[0].Handler != "health" {
		t.Errorf("handler = %q, want health", eps[0].Handler)
	}
}
