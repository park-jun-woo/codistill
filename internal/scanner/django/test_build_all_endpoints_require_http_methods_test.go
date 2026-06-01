//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestBuildAllEndpoints_RequireHTTPMethods — @require_http_methods(["POST","PUT"])가 POST·PUT 2건 방출
package django

import "testing"

func TestBuildAllEndpoints_RequireHTTPMethods(t *testing.T) {
	urls := mkFile(t, "app/urls.py", "app.urls",
		"from django.urls import path\nfrom .views import edit\nurlpatterns = [path('edit/', edit)]\n")
	views := mkFile(t, "app/views.py", "app.views",
		"from django.views.decorators.http import require_http_methods\n@require_http_methods([\"POST\", \"PUT\"])\ndef edit(request):\n    pass\n")

	eps := buildAllEndpoints([]fileInfo{urls, views})
	got := map[string]bool{}
	for _, ep := range eps {
		if ep.Path == "/edit/" {
			got[ep.Method] = true
		}
	}
	if !got["POST"] || !got["PUT"] || len(got) != 2 {
		t.Fatalf("expected POST and PUT for /edit/, got %v (all=%v)", got, eps)
	}
}
