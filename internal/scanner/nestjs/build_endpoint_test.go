//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildEndpoint_Basic 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildEndpoint_Basic(t *testing.T) {
	ci := controllerInfo{prefix: "users"}
	ep := endpointInfo{method: "GET", path: ":id", handler: "findOne", params: []scanner.Param{{Name: "id", Type: "string"}}}
	result := buildEndpoint("api", false, "", ci, ep)
	if result.Method != "GET" {
		t.Fatalf("expected GET, got %s", result.Method)
	}
	if result.Path != "/api/users/{id}" {
		t.Fatalf("expected /api/users/{id}, got %s", result.Path)
	}
}
