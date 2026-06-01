//ff:func feature=scan type=test control=sequence topic=actix
//ff:what Phase190 App 직속 .route("/health", web::get().to(h)) B형 추출 테스트
package actix

import "testing"

func TestExtractBuilderRoutes_AppDirectRoute(t *testing.T) {
	src := `fn cfg(app: App) -> App {
		app
			.service(web::resource("/sub").route(web::get().to(get_sub)))
			.route("/health", web::get().to(health_check))
	}`
	fi := aFi(t, src)
	routes := extractBuilderRoutes(fi, nil)
	r := findBuilderRoute(routes, "GET", "/health")
	if r == nil {
		t.Fatalf("expected GET /health, got %+v", routes)
	}
	if r.handler != "health_check" {
		t.Errorf("handler = %q, want health_check", r.handler)
	}
	if findBuilderRoute(routes, "GET", "/sub") == nil {
		t.Errorf("expected GET /sub (A-form unchanged), got %+v", routes)
	}
}
