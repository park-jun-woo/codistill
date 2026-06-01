//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what include(router.urls) 배선 + 커스텀 베이스 ViewSet CRUD 전개를 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_RouterURLsInclude(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets

class ModelCrudViewSet(viewsets.ModelViewSet):
    pass

class ProjectViewSet(ModelCrudViewSet):
    pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import ProjectViewSet

router = DefaultRouter()
router.register(r"projects", ProjectViewSet)

urlpatterns = [
    path('api/v1/', include(router.urls)),
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
		t.Logf("found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)
	}
	want := []string{
		"GET /api/v1/projects",
		"POST /api/v1/projects",
		"GET /api/v1/projects/{pk}",
		"PUT /api/v1/projects/{pk}",
		"PATCH /api/v1/projects/{pk}",
		"DELETE /api/v1/projects/{pk}",
	}
	for _, w := range want {
		if !found[w] {
			t.Errorf("missing router-wired CRUD endpoint %q", w)
		}
	}
	if len(result.Endpoints) != 6 {
		t.Errorf("expected exactly 6 CRUD endpoints (no flat duplicate), got %d", len(result.Endpoints))
	}
}
