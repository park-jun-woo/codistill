//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildSingleURLEntry_APIView 테스트
package django

import "testing"

func TestBuildSingleURLEntry_APIView(t *testing.T) {
	entry := urlEntry{pattern: "ping/", viewName: "PingView"}
	av := []apiviewInfo{{name: "PingView", methods: []string{"GET"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, nil, av, nil, map[string]serializerInfo{}, nil)
	if len(eps) == 0 {
		t.Fatal("expected apiview endpoints")
	}
}
