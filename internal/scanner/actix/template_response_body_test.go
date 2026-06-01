//ff:type feature=scan type=test topic=actix
//ff:what 응답 본문 스키마 추출 테스트용 소스 템플릿
package actix

var responseBodySource = `
use actix_web::{post, get, web, HttpResponse};

#[derive(Serialize)]
struct CreatedURL {
    shorturl: String,
    longurl: String,
    hits: i64,
}

#[post("/url")]
async fn create_url() -> HttpResponse {
    HttpResponse::Created().json(CreatedURL{shorturl: a, longurl: b, hits: c})
}

#[get("/url")]
async fn get_url() -> HttpResponse {
    let resp = CreatedURL{shorturl: a, longurl: b, hits: c};
    HttpResponse::Ok().json(resp)
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(create_url).service(get_url);
}
`
