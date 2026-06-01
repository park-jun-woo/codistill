//ff:func feature=scan type=test control=sequence topic=actix
//ff:what Phase190 회귀: .service()만 있고 App 직속 .route 없으면 B형 추가 없음
package actix

import "testing"

func TestExtractBuilderRoutes_ServiceOnlyNoDirectRoute(t *testing.T) {
	src := `fn cfg(app: App) -> App {
		app.service(web::scope("/api").service(web::resource("/x").route(web::get().to(x))))
	}`
	fi := aFi(t, src)
	routes := extractBuilderRoutes(fi, nil)
	if findBuilderRoute(routes, "GET", "/api/x") == nil {
		t.Fatalf("expected GET /api/x, got %+v", routes)
	}
	if findBuilderRoute(routes, "GET", "/") != nil {
		t.Errorf("unexpected spurious root route in %+v", routes)
	}
}
