//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what defaultRouteSegment가 RestApi/Api 접미사를 제거하고 소문자화하는지 검증한다
package flask

import "testing"

func TestDefaultRouteSegment(t *testing.T) {
	cases := map[string]string{
		"FooApi":        "foo",
		"ChartRestApi":  "chart",
		"DashboardView": "dashboardview",
		"Api":           "api",
	}
	for in, want := range cases {
		if got := defaultRouteSegment(in); got != want {
			t.Fatalf("defaultRouteSegment(%q)=%q, want %q", in, got, want)
		}
	}
}
