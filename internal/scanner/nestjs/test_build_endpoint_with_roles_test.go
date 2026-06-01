//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_WithRoles 테스트
package nestjs

import "testing"

func TestBuildEndpoint_WithRoles(t *testing.T) {
	ci := controllerInfo{prefix: "app"}
	ep := endpointInfo{
		method:     "GET",
		path:       "premium",
		handler:    "getPremium",
		middleware: []string{"JwtAuthGuard", "RolesGuard"},
		roles:      []string{"Role.premium"},
	}
	result := buildEndpoint("api", false, "", ci, ep)
	if len(result.Roles) != 1 || result.Roles[0] != "Role.premium" {
		t.Fatalf("expected Roles=[Role.premium], got %v", result.Roles)
	}
	if len(result.Middleware) != 2 {
		t.Fatalf("expected 2 Middleware, got %v", result.Middleware)
	}
}
