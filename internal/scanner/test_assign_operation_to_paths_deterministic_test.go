//ff:func feature=scan type=test control=sequence
//ff:what TestAssignOperationToPaths_Deterministic 테스트 (투입 순서를 바꿔도 richer op가 동일하게 잔존)
package scanner

import "testing"

func TestAssignOperationToPaths_Deterministic(t *testing.T) {
	rich := Endpoint{
		Method: "GET", Path: "/users", Handler: "rich", File: "a.rs", Line: 1,
		Responses: []Response{{Status: "200", TypeName: "User"}},
	}
	poor := Endpoint{Method: "GET", Path: "/users", Handler: "poor", File: "b.rs", Line: 2}

	keptHandler := func(first, second Endpoint) string {
		paths := map[string]map[string]any{}
		incumbent := map[string]map[string]Endpoint{}
		for _, ep := range []Endpoint{first, second} {
			if paths["/users"] == nil {
				paths["/users"] = map[string]any{}
			}
			assignOperationToPaths(paths, incumbent, "/users", ep, map[string]any{"h": ep.Handler})
		}
		op, _ := paths["/users"]["get"].(map[string]any)
		return op["h"].(string)
	}

	if got := keptHandler(rich, poor); got != "rich" {
		t.Fatalf("rich-first: expected rich kept, got %q", got)
	}
	if got := keptHandler(poor, rich); got != "rich" {
		t.Fatalf("poor-first: expected rich kept (order-independent), got %q", got)
	}
}
