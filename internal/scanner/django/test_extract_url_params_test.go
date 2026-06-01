//ff:func feature=scan type=test control=sequence topic=django
//ff:what Django URL 패턴에서 path parameter를 추출한다
package django

import "testing"

func TestExtractURLParams(t *testing.T) {
	params := extractURLParams("users/<int:pk>/posts/<slug:slug>/")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	if params[0].name != "pk" || params[0].converter != "int" {
		t.Errorf("param 0: expected pk (int), got %s (%s)", params[0].name, params[0].converter)
	}
	if params[1].name != "slug" || params[1].converter != "slug" {
		t.Errorf("param 1: expected slug (slug), got %s (%s)", params[1].name, params[1].converter)
	}

	rePathParams := extractURLParams("^articles/(?P<year>[0-9]{4})/$")
	if len(rePathParams) != 1 || rePathParams[0].name != "year" {
		t.Fatalf("expected re_path named group year, got %+v", rePathParams)
	}

	// Phase163: unnamed capture groups become positional {paramN} params,
	// matching the names djangoURLToOpenAPI synthesizes.
	unnamed := extractURLParams("^(\\d+)/(.*)")
	if len(unnamed) != 2 || unnamed[0].name != "param1" || unnamed[1].name != "param2" {
		t.Fatalf("expected positional param1/param2, got %+v", unnamed)
	}

	// Pure literal with escape/anchor yields no params.
	none := extractURLParams("^\\.well-known/jwks.json")
	if len(none) != 0 {
		t.Fatalf("expected no params for literal path, got %+v", none)
	}
}
