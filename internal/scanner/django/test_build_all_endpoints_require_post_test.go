//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestBuildAllEndpoints_RequirePost — @require_POST plain 뷰가 POST 엔드포인트로 방출되는지 검증
package django

import "testing"

func TestBuildAllEndpoints_RequirePost(t *testing.T) {
	urls := mkFile(t, "app/urls.py", "app.urls",
		"from django.urls import path\nfrom .views import webhook\nurlpatterns = [path('hook/', webhook)]\n")
	views := mkFile(t, "app/views.py", "app.views",
		"from django.views.decorators.http import require_POST\n@require_POST\ndef webhook(request):\n    pass\n")

	eps := buildAllEndpoints([]fileInfo{urls, views})
	var hook []string
	for _, ep := range eps {
		if ep.Path == "/hook/" {
			hook = append(hook, ep.Method)
		}
	}
	if len(hook) != 1 || hook[0] != "POST" {
		t.Fatalf("expected single POST endpoint for /hook/, got %v (all=%v)", hook, eps)
	}
}
