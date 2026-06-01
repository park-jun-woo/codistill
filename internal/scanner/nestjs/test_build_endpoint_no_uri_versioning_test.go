//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_NoURIVersioning 테스트
package nestjs

import "testing"

func TestBuildEndpoint_NoURIVersioning(t *testing.T) {
	ci := controllerInfo{prefix: "auth", version: "1"}
	ep := endpointInfo{method: "POST", path: "email/login", handler: "login"}
	result := buildEndpoint("api", false, "", ci, ep)
	if result.Path != "/api/auth/email/login" {
		t.Fatalf("expected /api/auth/email/login, got %s", result.Path)
	}
}
