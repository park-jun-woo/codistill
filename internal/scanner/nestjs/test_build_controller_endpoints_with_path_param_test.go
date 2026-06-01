//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildControllerEndpoints_WithPathParam 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildControllerEndpoints_WithPathParam(t *testing.T) {
	cwf := controllerWithFile{
		info: controllerInfo{
			prefix: "users",
			endpoints: []endpointInfo{
				{method: "GET", path: ":id", handler: "findOne", params: []scanner.Param{{Name: "id", Type: "string"}}},
			},
		},
		absFile: "/src/users.controller.ts",
	}
	eps, _ := buildControllerEndpoints("", false, "", cwf, "", 0)
	if len(eps) != 1 {
		t.Fatalf("expected 1, got %d", len(eps))
	}
	if eps[0].Path != "/users/{id}" {
		t.Fatalf("expected /users/{id}, got %s", eps[0].Path)
	}
}
