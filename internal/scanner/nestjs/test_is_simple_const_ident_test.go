//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestIsSimpleConstIdent 테스트 (식별자/멤버표현식/숫자 분기)
package nestjs

import "testing"

func TestIsSimpleConstIdent(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"HEALTH_CHECK_ROUTE", true},
		{"AUTH_LOGIN_PATH", true},
		{"$ref", true},
		{"_x", true},
		{"V2", true},
		{"", false},
		{"RouteKey.Asset", false},
		{"fn()", false},
		{"a[0]", false},
		{"`tpl`", false},
		{"200", false},
		{"2way", false},
		{"a b", false},
	}
	for _, c := range cases {
		if got := isSimpleConstIdent(c.in); got != c.want {
			t.Errorf("isSimpleConstIdent(%q): want %v got %v", c.in, c.want, got)
		}
	}
}
