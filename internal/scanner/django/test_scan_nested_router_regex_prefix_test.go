//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what nested router include의 regex-group 접두사 합성 + path param 승계를 E2E로 검증한다
package django

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScan_NestedRouterRegexPrefix(t *testing.T) {
	dir := t.TempDir()

	viewsSrc := `from rest_framework import viewsets

class ModelCrudViewSet(viewsets.ModelViewSet):
    pass

class ItemViewSet(ModelCrudViewSet):
    pass

class TicketViewSet(ModelCrudViewSet):
    pass
`
	os.WriteFile(filepath.Join(dir, "views.py"), []byte(viewsSrc), 0o644)

	urlsSrc := `from django.urls import re_path, include
from rest_framework.routers import DefaultRouter
from .views import ItemViewSet, TicketViewSet

orga_router = DefaultRouter()
orga_router.register(r"items", ItemViewSet)

event_router = DefaultRouter()
event_router.register(r"tickets", TicketViewSet)

urlpatterns = [
    re_path(r'^orgs/(?P<org>[^/]+)/', include(orga_router.urls)),
    re_path(r'^orgs/(?P<org>[^/]+)/events/(?P<event>[^/]+)/', include(event_router.urls)),
]
`
	os.WriteFile(filepath.Join(dir, "urls.py"), []byte(urlsSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	byKey := map[string]scanner.Endpoint{}
	for _, ep := range result.Endpoints {
		byKey[ep.Method+" "+ep.Path] = ep
		t.Logf("found: %s %s (%s)", ep.Method, ep.Path, ep.Handler)
	}

	// Single-level nested prefix: param "org" inherited; literal list + detail.
	wantSingle := []string{
		"GET /orgs/{org}/items",
		"POST /orgs/{org}/items",
		"GET /orgs/{org}/items/{pk}",
		"PUT /orgs/{org}/items/{pk}",
		"PATCH /orgs/{org}/items/{pk}",
		"DELETE /orgs/{org}/items/{pk}",
	}
	for _, w := range wantSingle {
		if _, ok := byKey[w]; !ok {
			t.Errorf("missing single-nested endpoint %q", w)
		}
	}

	// Two-level nested prefix: params "org" and "event" both inherited.
	wantDouble := []string{
		"GET /orgs/{org}/events/{event}/tickets",
		"GET /orgs/{org}/events/{event}/tickets/{pk}",
		"DELETE /orgs/{org}/events/{event}/tickets/{pk}",
	}
	for _, w := range wantDouble {
		if _, ok := byKey[w]; !ok {
			t.Errorf("missing double-nested endpoint %q", w)
		}
	}

	// Param inheritance order: prefix params precede {pk}.
	detail := byKey["GET /orgs/{org}/items/{pk}"]
	if detail.Request == nil {
		t.Fatalf("detail endpoint missing Request/path params")
	}
	names := []string{}
	for _, p := range detail.Request.PathParams {
		names = append(names, p.Name)
	}
	if len(names) != 2 || names[0] != "org" || names[1] != "pk" {
		t.Errorf("single-nested detail path params = %v, want [org pk]", names)
	}

	detail2 := byKey["GET /orgs/{org}/events/{event}/tickets/{pk}"]
	if detail2.Request == nil {
		t.Fatalf("double-nested detail endpoint missing Request/path params")
	}
	names2 := []string{}
	for _, p := range detail2.Request.PathParams {
		names2 = append(names2, p.Name)
	}
	if len(names2) != 3 || names2[0] != "org" || names2[1] != "event" || names2[2] != "pk" {
		t.Errorf("double-nested detail path params = %v, want [org event pk]", names2)
	}

	// List (non-detail) endpoint inherits prefix params even without {pk}.
	list := byKey["GET /orgs/{org}/items"]
	if list.Request == nil || len(list.Request.PathParams) != 1 || list.Request.PathParams[0].Name != "org" {
		t.Errorf("single-nested list path params want [org], got %+v", list.Request)
	}
}
