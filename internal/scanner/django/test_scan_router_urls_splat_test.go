//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what urlpatterns=[*router.urls] splat이 등록된 ViewSet들을 CRUD로 전개하는지 E2E 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_RouterURLsSplat(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets

class UserViewSet(viewsets.ModelViewSet):
    pass

class DocViewSet(viewsets.ModelViewSet):
    pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from rest_framework.routers import DefaultRouter
from .views import UserViewSet, DocViewSet

router = DefaultRouter()
router.register(r"users", UserViewSet)
router.register(r"docs", DocViewSet)

urlpatterns = [
    *router.urls,
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
		"GET /users", "GET /users/{pk}",
		"GET /docs", "GET /docs/{pk}",
	} {
		if !found[want] {
			t.Errorf("missing %q from *router.urls splat; got %v", want, found)
		}
	}
	// Two ModelViewSets => 2 x 6 CRUD endpoints.
	if len(result.Endpoints) != 12 {
		t.Errorf("expected 12 CRUD endpoints from 2 registered ViewSets, got %d", len(result.Endpoints))
	}
}
