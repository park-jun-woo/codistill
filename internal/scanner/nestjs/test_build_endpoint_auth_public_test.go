//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_AuthPublic 테스트
package nestjs

import "testing"

func TestBuildEndpoint_AuthPublic(t *testing.T) {
	ci := controllerInfo{prefix: "auth"}
	ep := endpointInfo{
		method:    "POST",
		path:      "login",
		handler:   "signIn",
		authLevel: "public",
	}
	result := buildEndpoint("api", false, "", ci, ep)
	if result.AuthLevel != "public" {
		t.Fatalf("expected AuthLevel=public, got %q", result.AuthLevel)
	}
}
