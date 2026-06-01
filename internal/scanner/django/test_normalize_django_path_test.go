//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what Django URL 정규식을 OpenAPI path 문자열+param 목록으로 완전 정규화하는 동작을 검증한다
package django

import (
	"strings"
	"testing"
)

func TestNormalizeDjangoPath(t *testing.T) {
	tests := []struct {
		input      string
		wantPath   string
		wantParams string // comma-joined param names
	}{
		{"^403-csrf-failure/$", "403-csrf-failure/", ""},
		{"^(\\d+)/(.*)", "{param1}/{param2}", "param1,param2"},
		{"^\\.well-known/jwks.json", ".well-known/jwks.json", ""},
		{"^articles/(?P<year>[0-9]{4})/$", "articles/{year}/", "year"},
		{"users/<int:pk>/", "users/{pk}/", "pk"},
		{"^(?:v1/)?items/(?P<pk>\\d+)/", "items/{pk}/", "pk"},
	}
	for _, tt := range tests {
		gotPath, gotParams := normalizeDjangoPath(tt.input)
		var names []string
		for _, p := range gotParams {
			names = append(names, p.name)
		}
		gotNames := strings.Join(names, ",")
		if gotPath != tt.wantPath || gotNames != tt.wantParams {
			t.Errorf("normalizeDjangoPath(%q) = (%q, %q), want (%q, %q)",
				tt.input, gotPath, gotNames, tt.wantPath, tt.wantParams)
		}
	}
}
