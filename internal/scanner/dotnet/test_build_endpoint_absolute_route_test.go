//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestBuildEndpointAbsoluteRoute -- ~/ · / 절대 라우트가 prefix와 합성되지 않음을 검증
package dotnet

import "testing"

func TestBuildEndpointAbsoluteRoute(t *testing.T) {
	cases := []struct {
		prefix string
		path   string
		want   string
	}{
		{"apps", "~/api/v1/stores", "/api/v1/stores"},
		{"apps", "/apps/{appId}/pos", "/apps/{appId}/pos"},
		{"apps", "/", "/"},
		{"stores", "{storeId}/index", "/stores/{storeId}/index"},
		{"", "/login", "/login"},
	}
	for _, c := range cases {
		ci := controllerInfo{prefix: c.prefix, className: "C"}
		ep := endpointInfo{method: "GET", path: c.path, handler: "H"}
		got := buildEndpoint(ci, ep)
		if got.Path != c.want {
			t.Errorf("prefix=%q path=%q => %q, want %q", c.prefix, c.path, got.Path, c.want)
		}
	}
}
