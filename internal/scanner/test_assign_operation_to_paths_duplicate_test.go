//ff:func feature=scan type=test control=sequence
//ff:what TestAssignOperationToPaths_Duplicate 테스트 (충돌 시 preferEndpoint 결정적 보존 + 손실 집계)
package scanner

import "testing"

func TestAssignOperationToPaths_Duplicate(t *testing.T) {
	// incumbent get은 타입 정보가 풍부(richer)한 핸들러, 새 후보는 빈약(poorer).
	rich := Endpoint{
		Method: "GET", Handler: "richList", File: "a.rs", Line: 1,
		Responses: []Response{{Status: "200", TypeName: "User"}},
	}
	paths := map[string]map[string]any{"/users": {"get": map[string]any{"old": true}}}
	incumbent := map[string]map[string]Endpoint{"/users": {"get": rich}}

	poor := Endpoint{Method: "GET", Handler: "poorList", File: "b.rs", Line: 2}
	conflicts := assignOperationToPaths(paths, incumbent, "/users", poor, map[string]any{"new": true})

	// richer incumbent가 보존되어야 하므로 op는 덮이지 않는다.
	got, _ := paths["/users"]["get"].(map[string]any)
	if got["old"] != true {
		t.Fatalf("expected richer incumbent preserved, got %+v", paths["/users"]["get"])
	}
	if len(conflicts) != 1 {
		t.Fatalf("expected 1 conflict recorded, got %d", len(conflicts))
	}
	if conflicts[0].DropHandler != "poorList" || conflicts[0].KeptHandler != "richList" {
		t.Fatalf("unexpected conflict record: %+v", conflicts[0])
	}
}
