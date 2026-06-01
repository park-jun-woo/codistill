//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what 클래스-레벨 [Route("~/...")]·[Route("/...")] 절대 prefix 선두 마커 strip 검증(상대 prefix 회귀 가드)
package dotnet

import "testing"

func TestExtractClassRouteAbsolutePrefix(t *testing.T) {
	cases := []struct {
		route string
		name  string
		want  string
	}{
		{`[Route("~/{cryptoCode}/[controller]/")]`, "UILNURL", "{cryptoCode}/uilnurl/"},
		{`[Route("~/api/v1/stores")]`, "Stores", "api/v1/stores"},
		{`[Route("/apps")]`, "Apps", "apps"},
		{`[Route("stores")]`, "Stores", "stores"},             // 상대 회귀 가드
		{`[Route("api/[controller]")]`, "Users", "api/users"}, // 상대 회귀 가드
	}
	for _, c := range cases {
		root, src := parseCS(t, c.route+` class C {}`)
		cls := findAllByType(root, "class_declaration")[0]
		got := extractClassRoute(cls, src, c.name)
		if got != c.want {
			t.Errorf("route=%q => %q, want %q", c.route, got, c.want)
		}
	}
}
