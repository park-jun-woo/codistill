//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildSingleURLEntry_Plain 테스트
package django

import "testing"

func TestBuildSingleURLEntry_Plain(t *testing.T) {

	entry := urlEntry{pattern: "misc/", viewName: "UnknownView"}
	eps := buildSingleURLEntryEndpoints(entry, nil, nil, nil, map[string]serializerInfo{}, nil)
	if len(eps) == 0 {
		t.Fatal("expected plain endpoints fallback")
	}
}
