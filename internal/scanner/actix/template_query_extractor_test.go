//ff:type feature=scan type=test topic=actix
//ff:what Query extractor 테스트용 소스 템플릿
package actix

var queryExtractorSource = `
use actix_web::{get, web, HttpResponse};

#[derive(Deserialize)]
struct ListParams {
    page: Option<u32>,
    limit: Option<u32>,
    search: Option<String>,
}

#[get("/items")]
async fn list_items(query: web::Query<ListParams>) -> HttpResponse {
    HttpResponse::Ok().json(items)
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(list_items);
}
`
