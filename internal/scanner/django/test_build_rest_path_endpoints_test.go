//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestBuildRestPathEndpoints 테스트
package django

import "testing"

func TestBuildRestPathEndpoints(t *testing.T) {
	entry := urlEntry{
		pattern:     "realm",
		methodViews: map[string]string{"GET": "get_realm", "PATCH": "views.update_realm"},
	}
	eps := buildRestPathEndpoints(entry)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	byMethod := map[string]string{}
	for _, ep := range eps {
		if ep.Path != "/realm" {
			t.Errorf("path: %q", ep.Path)
		}
		byMethod[ep.Method] = ep.Handler
	}
	if byMethod["GET"] != "get_realm" {
		t.Errorf("GET handler: %q", byMethod["GET"])
	}
	if byMethod["PATCH"] != "update_realm" {
		t.Errorf("PATCH handler: %q", byMethod["PATCH"])
	}
}
