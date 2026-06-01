//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_AuthRequired 테스트
package nestjs

import "testing"

func TestBuildEndpoint_AuthRequired(t *testing.T) {
	ci := controllerInfo{prefix: "users"}
	ep := endpointInfo{
		method:    "GET",
		path:      "me",
		handler:   "getMe",
		authLevel: "auth_required",
	}
	result := buildEndpoint("api", false, "", ci, ep)
	if result.AuthLevel != "auth_required" {
		t.Fatalf("expected AuthLevel=auth_required, got %q", result.AuthLevel)
	}
}
