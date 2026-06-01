//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what include([*router.urls, path(...)]) 인라인 리스트가 접두사 하에 splat+path를 함께 전개하는지 E2E 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_InlineIncludeSplat(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets
from rest_framework.views import APIView

class UserViewSet(viewsets.ModelViewSet):
    pass

class HealthView(APIView):
    def get(self, request):
        pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import UserViewSet, HealthView

router = DefaultRouter()
router.register(r"users", UserViewSet)

urlpatterns = [
    path("api/", include([
        *router.urls,
        path("health/", HealthView.as_view()),
    ])),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urlsSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	for _, want := range []string{
		"GET /api/users", "GET /api/users/{pk}",
		"GET /api/health/",
	} {
		if !found[want] {
			t.Errorf("missing %q from inline include splat; got %v", want, found)
		}
	}
	// The router is wired through the include prefix; it must not also be expanded
	// as a bare (prefix-less) root by the flat router pass.
	if found["GET /users"] {
		t.Errorf("unexpected prefix-less /users; wired router duplicated, got %v", found)
	}
}
