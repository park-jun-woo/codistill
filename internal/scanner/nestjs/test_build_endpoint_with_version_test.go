//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_WithVersion 테스트
package nestjs

import "testing"

func TestBuildEndpoint_WithVersion(t *testing.T) {
	ci := controllerInfo{prefix: "auth", version: "1"}
	ep := endpointInfo{method: "POST", path: "email/login", handler: "login"}
	result := buildEndpoint("api", true, "", ci, ep)
	if result.Path != "/api/v1/auth/email/login" {
		t.Fatalf("expected /api/v1/auth/email/login, got %s", result.Path)
	}
}
