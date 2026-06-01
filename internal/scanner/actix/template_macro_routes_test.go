//ff:type feature=scan type=test topic=actix
//ff:what proc-macro 라우트 테스트용 소스 템플릿
package actix

var macroRoutesSource = `
use actix_web::{get, post, put, delete, web, HttpResponse};

#[derive(Deserialize)]
struct CreateUserRequest {
    name: String,
    email: String,
    #[serde(default)]
    age: Option<i32>,
}

#[get("/users/{id}")]
async fn get_user(path: web::Path<i64>) -> HttpResponse {
    HttpResponse::Ok().json(user)
}

#[post("/users")]
async fn create_user(body: web::Json<CreateUserRequest>) -> HttpResponse {
    HttpResponse::Created().json(user)
}

#[put("/users/{id}")]
async fn update_user(
    path: web::Path<i64>,
    body: web::Json<CreateUserRequest>,
) -> HttpResponse {
    HttpResponse::Ok().json(user)
}

#[delete("/users/{id}")]
async fn delete_user(path: web::Path<i64>) -> HttpResponse {
    HttpResponse::NoContent().finish()
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(get_user)
        .service(create_user)
        .service(update_user)
        .service(delete_user);
}
`
