//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestFirstPathSegment 테스트
package scanner

import "testing"

func TestFirstPathSegment(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/api/v1/categories/find/{id}", "categories"},
		{"/api/v1/users", "users"},
		{"/api/v2/products/:id", "products"},
		{"/health", "health"},
		{"/{param}", ""},
		{"/api/v1/{id}", ""},
		{"", ""},
		{"/api", ""},
		{"/api/v1", ""},
		{"/admin/buildings/:buildingId", "admin"},
		// Phase145: skip unresolved `${PREFIX}` template fragment.
		{"/${PREFIX}/{modelId}/records", "records"},
		{"/${PREFIX}", ""},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := firstPathSegment(tt.path)
			if got != tt.want {
				t.Errorf("firstPathSegment(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
