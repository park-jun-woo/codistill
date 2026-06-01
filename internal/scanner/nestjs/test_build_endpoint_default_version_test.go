//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_DefaultVersionFallback 테스트
package nestjs

import "testing"

func TestBuildEndpoint_DefaultVersionFallback(t *testing.T) {
	// Novu case: @Controller('/activity') has no version; app-level
	// defaultVersion '1' should be applied → /v1/activity/requests.
	ci := controllerInfo{prefix: "/activity"}
	ep := endpointInfo{method: "GET", path: "requests", handler: "getActivity"}
	result := buildEndpoint("", true, "1", ci, ep)
	if result.Path != "/v1/activity/requests" {
		t.Fatalf("expected /v1/activity/requests, got %s", result.Path)
	}
}
