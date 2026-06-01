//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestAssignSchemaSection_Meta 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestAssignSchemaSection_Meta(t *testing.T) {
	pairs, src := schemaPairs(t, `{
  operationId: "healthCheck",
  summary: "Health",
  description: "Application health check endpoint.",
  tags: ["system", "ops"]
}`)
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	for _, p := range pairs {
		assignSchemaSection(si, p, src)
	}
	if si.OperationID != "healthCheck" {
		t.Fatalf("operationId: got %q", si.OperationID)
	}
	if si.Summary != "Health" {
		t.Fatalf("summary: got %q", si.Summary)
	}
	if si.Description != "Application health check endpoint." {
		t.Fatalf("description: got %q", si.Description)
	}
	if len(si.Tags) != 2 || si.Tags[0] != "system" || si.Tags[1] != "ops" {
		t.Fatalf("tags: got %v", si.Tags)
	}
}
