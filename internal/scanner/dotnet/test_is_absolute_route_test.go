//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestIsAbsoluteRoute -- ~/ · / 절대 라우트 판정 및 선두 마커 strip 테스트
package dotnet

import "testing"

func TestIsAbsoluteRoute(t *testing.T) {
	cases := []struct {
		in       string
		want     string
		wantAbs  bool
	}{
		{"~/api/v1/stores", "api/v1/stores", true},
		{"/apps/{appId}/crowdfund/form", "apps/{appId}/crowdfund/form", true},
		{"/", "", true},
		{"~", "", true},
		{"{storeId}/index", "{storeId}/index", false},
		{"api/v1/stores", "api/v1/stores", false},
	}
	for _, c := range cases {
		got, abs := isAbsoluteRoute(c.in)
		if got != c.want || abs != c.wantAbs {
			t.Errorf("isAbsoluteRoute(%q) = (%q, %v), want (%q, %v)", c.in, got, abs, c.want, c.wantAbs)
		}
	}
}
