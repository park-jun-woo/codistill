//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what Django URL 패턴을 OpenAPI 형식으로 변환한다
package django

import "testing"

func TestDjangoURLToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"users/<int:pk>/", "users/{pk}/"},
		{"users/<int:pk>/posts/<int:post_id>/", "users/{pk}/posts/{post_id}/"},
		{"users/<str:slug>/", "users/{slug}/"},
		{"users/<uuid:id>/", "users/{id}/"},
		{"health/", "health/"},
		{"", ""},
		{"^articles/(?P<year>[0-9]{4})/$", "articles/{year}/"},
		{"^users/(?P<slug>[-\\w]+)/$", "users/{slug}/"},
		// Phase163: anchors removed at every position, unnamed groups, escapes.
		{"^403-csrf-failure/$", "403-csrf-failure/"},
		{"/api/^control/dashboard", "/api/control/dashboard"},
		{"^(\\d+)/(.*)", "{param1}/{param2}"},
		{"^\\.well-known/jwks.json", ".well-known/jwks.json"},
		{"^([^/]*)/(\\d*)/edit/$", "{param1}/{param2}/edit/"},
		{"users/{pk}/", "users/{pk}/"},
		{"^graphql/$", "graphql/"},
	}

	for _, tt := range tests {
		got := djangoURLToOpenAPI(tt.input)
		if got != tt.want {
			t.Errorf("djangoURLToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
