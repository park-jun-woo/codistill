//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestMacroRoute_RegisteredEmitted — .service(h) 등록된 매크로 라우트는 정상 추출한다
package actix

import "testing"

func TestMacroRoute_RegisteredEmitted(t *testing.T) {
	src := `
use actix_web::{get, web, App, HttpResponse};

#[get("/x")]
async fn h() -> HttpResponse {
    HttpResponse::Ok().finish()
}

pub fn main() {
    App::new().service(h);
}
`
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d: %+v", len(result.Endpoints), result.Endpoints)
	}
	ep := result.Endpoints[0]
	if ep.Method != "GET" || ep.Path != "/x" {
		t.Fatalf("expected GET /x, got %s %s", ep.Method, ep.Path)
	}
}
