//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestMacroRoute_CommentedDropped — 주석 처리된 .service(h)는 line_comment라 등록 집합에서 제외되어 매크로 라우트가 추출되지 않는다
package actix

import "testing"

func TestMacroRoute_CommentedDropped(t *testing.T) {
	src := `
use actix_web::{get, web, App, HttpResponse};

#[get("/x")]
async fn h() -> HttpResponse {
    HttpResponse::Ok().finish()
}

pub fn main() {
    App::new()
        //.service(h)
        ;
}
`
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints (commented registration), got %d: %+v", len(result.Endpoints), result.Endpoints)
	}
}
