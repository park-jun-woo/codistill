//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what urlpatterns = router.urls 배선(접두사 없음) CRUD 전개를 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_RouterURLsAssignment(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets

class UserViewSet(viewsets.ModelViewSet):
    pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from rest_framework.routers import DefaultRouter
from .views import UserViewSet

router = DefaultRouter()
router.register(r"users", UserViewSet)

urlpatterns = router.urls
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
	if !found["GET /users"] || !found["GET /users/{pk}"] {
		t.Errorf("missing router.urls-assigned CRUD endpoints, got %v", found)
	}
	if len(result.Endpoints) != 6 {
		t.Errorf("expected 6 CRUD endpoints, got %d", len(result.Endpoints))
	}
}
