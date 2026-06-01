//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what Express 경로 → OpenAPI 경로 변환 테스트: :id → {id}
package express

import "testing"

func TestExpressPathToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/users/:id", "/users/{id}"},
		{"/users/:userId/posts/:postId", "/users/{userId}/posts/{postId}"},
		{"/users", "/users"},
		{"/:id", "/{id}"},
		{"/files/:name(\\d+)", "/files/{name}"},
		{"/api/badge/:id/avg-response/:duration?", "/api/badge/{id}/avg-response/{duration}"},
	}
	for _, tt := range tests {
		got := expressPathToOpenAPI(tt.input)
		if got != tt.want {
			t.Errorf("expressPathToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
