//ff:func feature=scan type=test control=sequence topic=actix
//ff:what Phase190 scope 직속 .route("/ping", ...) B형 prefix 합성 테스트
package actix

import "testing"

func TestExtractBuilderRoutes_ScopeDirectRoute(t *testing.T) {
	src := `fn cfg(app: App) -> App {
		app.service(web::scope("/api").route("/ping", web::get().to(ping)))
	}`
	fi := aFi(t, src)
	routes := extractBuilderRoutes(fi, nil)
	r := findBuilderRoute(routes, "GET", "/api/ping")
	if r == nil {
		t.Fatalf("expected GET /api/ping (scope prefix composed), got %+v", routes)
	}
	if r.handler != "ping" {
		t.Errorf("handler = %q, want ping", r.handler)
	}
}
