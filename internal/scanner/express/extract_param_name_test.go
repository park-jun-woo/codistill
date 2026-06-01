//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what extractParamName: 콜론없음/일반파라미터/정규식파라미터 분기
package express

import "testing"

func TestExtractParamName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"users", ""},          // no colon prefix
		{":id", "id"},          // plain param
		{":id(\\d+)", "id"},    // regex constraint -> strip from "("
		{":duration?", "duration"}, // optional param -> strip trailing "?"
	}
	for _, c := range cases {
		if got := extractParamName(c.in); got != c.want {
			t.Errorf("extractParamName(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
