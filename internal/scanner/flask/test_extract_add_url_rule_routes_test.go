//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what extractAddURLRuleRoutes의 prefix합성·!앱루트·tuple메서드·기본GET을 검증한다
package flask

import "testing"

func TestExtractAddURLRuleRoutes_PrefixAndAppRoot(t *testing.T) {
	fi := flaskFile(t, `_bp.add_url_rule('!/admin/admins', 'admins', RHAdmins, methods=('GET', 'POST'))
_bp.add_url_rule('/profile', 'profile', RHProfile)
`)
	bpPrefixes := blueprintPrefix{"_bp": "/users"}
	routes := extractAddURLRuleRoutes([]fileInfo{fi}, bpPrefixes)

	// '!' app-root: GET+POST at /admin/admins (prefix ignored). '/profile' -> /users/profile GET.
	if len(routes) != 3 {
		t.Fatalf("expected 3 routes, got %d: %+v", len(routes), routes)
	}
	got := map[string]string{}
	for _, r := range routes {
		got[r.method+" "+r.path] = r.handler
	}
	if h, ok := got["GET /admin/admins"]; !ok || h != "RHAdmins" {
		t.Errorf("missing GET /admin/admins -> RHAdmins: %+v", routes)
	}
	if _, ok := got["POST /admin/admins"]; !ok {
		t.Errorf("missing POST /admin/admins: %+v", routes)
	}
	if h, ok := got["GET /users/profile"]; !ok || h != "RHProfile" {
		t.Errorf("missing GET /users/profile -> RHProfile (default GET + prefix): %+v", routes)
	}
}
