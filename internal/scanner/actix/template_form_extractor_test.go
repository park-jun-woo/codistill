//ff:type feature=scan type=test topic=actix
//ff:what Form extractor 테스트용 소스 템플릿
package actix

var formExtractorSource = `
use actix_web::{post, web, HttpResponse};

#[derive(Deserialize)]
struct LoginForm {
    username: String,
    password: String,
}

#[post("/login")]
async fn login(form: web::Form<LoginForm>) -> HttpResponse {
    HttpResponse::Ok().json(result)
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(login);
}
`
