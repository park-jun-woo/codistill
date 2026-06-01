//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildSingleURLEntry_ViewSet 테스트
package django

import "testing"

func TestBuildSingleURLEntry_ViewSet(t *testing.T) {
	entry := urlEntry{pattern: "users/", viewName: "UserViewSet"}
	vs := []viewsetInfo{{name: "UserViewSet", parents: []string{"ModelViewSet"}, file: "v.py"}}
	eps := buildSingleURLEntryEndpoints(entry, vs, nil, nil, map[string]serializerInfo{}, nil)
	if len(eps) == 0 {
		t.Fatal("expected viewset endpoints")
	}
}
