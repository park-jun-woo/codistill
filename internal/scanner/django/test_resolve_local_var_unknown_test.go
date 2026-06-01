//ff:func feature=scan type=test control=sequence topic=django
//ff:what resolveLocalVarIncludes — 미정의 로컬변수 include는 라우트를 만들지 않고 drop되는지 검증
package django

import "testing"

func TestResolveLocalVarIncludesUnknownVarDropped(t *testing.T) {
	got := expandFile(t, `urlpatterns = [path("api/v1/", include(missing_urls))]`)
	if len(got) != 0 {
		t.Errorf("unknown local var should yield no routes, got %v", got)
	}
}
