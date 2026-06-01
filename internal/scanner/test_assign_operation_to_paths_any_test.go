//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestAssignOperationToPaths_Any 테스트
package scanner

import "testing"

func TestAssignOperationToPaths_Any(t *testing.T) {
	paths := map[string]map[string]any{"/x": {}}
	incumbent := map[string]map[string]Endpoint{}
	ep := Endpoint{Method: "any", Handler: "h"}
	assignOperationToPaths(paths, incumbent, "/x", ep, map[string]any{"k": 1})

	for _, m := range []string{"get", "post", "put", "patch", "delete"} {
		if paths["/x"][m] == nil {
			t.Errorf("expected method %q assigned for any", m)
		}
	}
}
