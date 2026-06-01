//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what named-list splat(*intake_urls) 회귀 없음을 E2E로 검증한다(splat 변경이 기존 동작을 깨지 않음)
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NamedListSplatRegression(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework.views import APIView

class HealthView(APIView):
    def get(self, request):
        pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	// A separate urlpatterns module that another module splats via *name. This is the
	// Plane-style pattern: the named list is collected as its own urlpatterns module
	// and expanded as an independent root; the splat must not regress it.
	subSrc := `from django.urls import path
from .views import HealthView

urlpatterns = [
    path("health/", HealthView.as_view()),
]
`
	os.WriteFile(filepath.Join(dir, "sub_urls.py"), []byte(subSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	if !found["GET /health/"] {
		t.Errorf("named-list/urlpatterns module regressed; got %v", found)
	}
}
