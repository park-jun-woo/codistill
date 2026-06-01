//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 함수명이 알려진 Django/DRF 뷰 래퍼 데코레이터인지 판별한다
package django

import "strings"

// knownViewWrappers is the set of common Django/DRF view-wrapping decorators
// that, when applied directly as a call around the real view in a URL pattern
// (e.g. staff_member_required(my_view)), take the wrapped view as their first
// positional argument. The wrapper name itself must not become the view name.
//
// Decorators that are configured then curried (e.g. cache_page(60)(view),
// permission_required("perm")(view), api_view([...]) ) put their view in a
// nested call whose callee is itself a call node; those are not matched here
// because resolveCallArg only reaches this branch for an identifier/attribute
// callee, and the unwrap step additionally requires the inner argument to
// resolve to a plausible view (see resolveCallArg) before adopting it.
var knownViewWrappers = map[string]bool{
	"staff_member_required":     true,
	"login_required":            true,
	"csrf_exempt":               true,
	"csrf_protect":              true,
	"ensure_csrf_cookie":        true,
	"never_cache":               true,
	"gzip_page":                 true,
	"xframe_options_exempt":     true,
	"xframe_options_deny":       true,
	"xframe_options_sameorigin": true,
}

// isViewWrapper reports whether name is a known view-wrapping decorator whose
// first positional argument is the real view. The "require_*" family without
// arguments (require_GET, require_POST, require_safe) wraps the view directly;
// argument-taking forms such as require_http_methods([...]) are curried and so
// never reach this branch with the view as their first positional argument, so
// the conservative unwrap guard in resolveCallArg keeps them safe.
func isViewWrapper(name string) bool {
	if knownViewWrappers[name] {
		return true
	}
	if strings.HasPrefix(name, "require_") {
		return true
	}
	return false
}
