//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestMacroRoute_UnregisteredDropped — .service() 등록 없는 매크로 라우트는 추출하지 않는다
package actix

import "testing"

func TestMacroRoute_UnregisteredDropped(t *testing.T) {
	src := `
use actix_web::{get, HttpResponse};

#[get("/x")]
async fn h() -> HttpResponse {
    HttpResponse::Ok().finish()
}
`
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints (dead route), got %d: %+v", len(result.Endpoints), result.Endpoints)
	}
}
