//ff:func feature=scan type=test control=sequence topic=flask
//ff:what resolveRegPrefix가 blueprint 우선·namespace 차선·미지정 빈문자열을 해석하는지 검증한다
package flask

import "testing"

func TestResolveRegPrefix(t *testing.T) {
	bp := blueprintPrefix{"users_bp": "/api/users"}
	ns := namespacePrefix{"inner_api_ns": "/inner/api"}

	if got := resolveRegPrefix("users_bp", bp, ns); got != "/api/users" {
		t.Errorf("blueprint: got %q", got)
	}
	if got := resolveRegPrefix("inner_api_ns", bp, ns); got != "/inner/api" {
		t.Errorf("namespace: got %q", got)
	}
	if got := resolveRegPrefix("api", bp, ns); got != "" {
		t.Errorf("unknown: got %q want empty", got)
	}
}
