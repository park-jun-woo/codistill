//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestBuildControllerEndpoints_Basic 테스트
package nestjs

import "testing"

func TestBuildControllerEndpoints_Basic(t *testing.T) {
	cwf := controllerWithFile{
		info: controllerInfo{
			prefix: "users",
			endpoints: []endpointInfo{
				{method: "GET", handler: "findAll"},
				{method: "POST", handler: "create", bodyType: "Dto", statusCode: 201},
			},
			imports: map[string]string{"Dto": "./dto"},
		},
		absFile: "/src/users.controller.ts",
	}
	eps, reqs := buildControllerEndpoints("api", false, "", cwf, "", 0)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Path != "/api/users" {
		t.Fatalf("expected /api/users, got %s", eps[0].Path)
	}
	found := false
	for _, r := range reqs {
		if r.typeName == "Dto" && r.isBody {
			found = true
		}
	}
	if !found {
		t.Fatal("expected body DTO request")
	}
}
