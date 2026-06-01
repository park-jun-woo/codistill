//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestNormalizePathRegex 테스트
package quarkus

import "testing"

func TestNormalizePathRegex(t *testing.T) {
	cases := map[string]string{
		"/{id:[0-9]+}":              "/{id}",
		"/{subscriptionId:" + "UUID_PATTERN" + "}/{unitType}": "/{subscriptionId}/{unitType}",
		"/plain/{name}":             "/plain/{name}",
		"/no/braces":                "/no/braces",
		"/{a:x}/{b}/{c:y}":          "/{a}/{b}/{c}",
	}
	for in, want := range cases {
		if got := normalizePathRegex(in); got != want {
			t.Errorf("normalizePathRegex(%q)=%q want %q", in, got, want)
		}
	}
}
