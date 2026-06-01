//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_AuthEmpty 테스트
package nestjs

import "testing"

func TestBuildEndpoint_AuthEmpty(t *testing.T) {
	ci := controllerInfo{prefix: "app"}
	ep := endpointInfo{
		method:  "GET",
		path:    "status",
		handler: "getStatus",
	}
	result := buildEndpoint("", false, "", ci, ep)
	if result.AuthLevel != "" {
		t.Fatalf("expected empty AuthLevel, got %q", result.AuthLevel)
	}
}
