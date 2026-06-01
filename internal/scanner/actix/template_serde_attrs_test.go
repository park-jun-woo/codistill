//ff:type feature=scan type=test topic=actix
//ff:what serde 어트리뷰트 테스트용 소스 템플릿
package actix

var serdeAttrsSource = `
use actix_web::{post, web, HttpResponse};

#[derive(Deserialize)]
struct UpdateRequest {
    #[serde(rename = "userName")]
    name: String,
    #[serde(default)]
    bio: Option<String>,
    #[serde(skip)]
    internal_id: String,
}

#[post("/update")]
async fn update(body: web::Json<UpdateRequest>) -> HttpResponse {
    HttpResponse::Ok().json(result)
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(update);
}
`
