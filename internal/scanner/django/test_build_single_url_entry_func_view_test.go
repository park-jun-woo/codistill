//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildSingleURLEntry_FuncView 테스트
package django

import "testing"

func TestBuildSingleURLEntry_FuncView(t *testing.T) {
	entry := urlEntry{pattern: "health/", viewName: "health"}
	fv := []funcViewInfo{{name: "health", methods: []string{"GET"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, nil, nil, fv, map[string]serializerInfo{}, nil)
	if len(eps) == 0 {
		t.Fatal("expected funcview endpoints")
	}
}
