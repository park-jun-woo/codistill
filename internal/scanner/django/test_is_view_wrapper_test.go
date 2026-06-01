//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestIsViewWrapper 테스트 (Phase165)
package django

import "testing"

func TestIsViewWrapper(t *testing.T) {
	for _, name := range []string{"staff_member_required", "login_required", "csrf_exempt", "require_POST", "require_http_methods"} {
		if !isViewWrapper(name) {
			t.Errorf("isViewWrapper(%q) = false, want true", name)
		}
	}
	for _, name := range []string{"include", "some_unknown_helper", "my_view", ""} {
		if isViewWrapper(name) {
			t.Errorf("isViewWrapper(%q) = true, want false", name)
		}
	}
}
