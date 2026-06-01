//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildOneEndpoint_Meta 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildOneEndpoint_Meta(t *testing.T) {
	src := `import Fastify from "fastify";
const app = Fastify();
app.get("/", {
  schema: {
    operationId: "healthCheck",
    summary: "Health",
    description: "Application health check endpoint.",
    tags: ["system"]
  }
}, async (req, reply) => reply.send());
`
	fi := mustParse(t, []byte(src))
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	ep := buildOneEndpoint("GET", "/api", routes[0], "indexController.ts", nil, fi.Src, map[string]*sitter.Node{})
	if ep.OperationID != "healthCheck" {
		t.Fatalf("operationId: got %q", ep.OperationID)
	}
	if ep.Summary != "Health" {
		t.Fatalf("summary: got %q", ep.Summary)
	}
	if ep.Description != "Application health check endpoint." {
		t.Fatalf("description: got %q", ep.Description)
	}
	if len(ep.Tags) != 1 || ep.Tags[0] != "system" {
		t.Fatalf("tags: got %v", ep.Tags)
	}
}
