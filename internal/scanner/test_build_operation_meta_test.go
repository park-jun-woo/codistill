//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperation_Meta 테스트
package scanner

import "testing"

func TestBuildOperation_Meta(t *testing.T) {
	ep := Endpoint{
		Method:      "GET",
		Path:        "/api",
		Handler:     "(anonymous)",
		Summary:     "Health",
		Description: "Application health check endpoint.",
		Tags:        []string{"system"},
	}
	op := buildOperation(ep, map[string]any{})
	if op["summary"] != "Health" {
		t.Fatalf("expected summary, got %v", op["summary"])
	}
	if op["description"] != "Application health check endpoint." {
		t.Fatalf("expected description, got %v", op["description"])
	}
	tags, ok := op["tags"].([]any)
	if !ok || len(tags) != 1 || tags[0] != "system" {
		t.Fatalf("expected tags [system], got %v", op["tags"])
	}
}
