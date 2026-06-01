//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_WithBody 테스트
package nestjs

import "testing"

func TestBuildEndpoint_WithBody(t *testing.T) {
	ci := controllerInfo{prefix: "users"}
	ep := endpointInfo{method: "POST", handler: "create", bodyType: "CreateUserDto", statusCode: 201}
	result := buildEndpoint("", false, "", ci, ep)
	if result.Request == nil || result.Request.Body == nil {
		t.Fatal("expected request body")
	}
	if len(result.Responses) == 0 || result.Responses[0].Status != "201" {
		t.Fatal("expected 201 status")
	}
}
