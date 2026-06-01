//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_ControllerVersionPrecedence 테스트
package nestjs

import "testing"

func TestBuildEndpoint_ControllerVersionPrecedence(t *testing.T) {
	// Hoppscotch regression guard: controller version '2' wins over the
	// app-level defaultVersion '1' → /v2/auth/...
	ci := controllerInfo{prefix: "auth", version: "2"}
	ep := endpointInfo{method: "GET", path: "verify", handler: "verify"}
	result := buildEndpoint("", true, "1", ci, ep)
	if result.Path != "/v2/auth/verify" {
		t.Fatalf("expected /v2/auth/verify, got %s", result.Path)
	}
}
