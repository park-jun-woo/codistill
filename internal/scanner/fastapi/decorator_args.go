//ff:type feature=scan type=model topic=fastapi
//ff:what 라우트 데코레이터에서 추출한 인자 묶음 구조체
package fastapi

// decoratorArgs holds arguments parsed from a route decorator like
// @app.get("/x", status_code=200, response_model=Foo, include_in_schema=False).
type decoratorArgs struct {
	path            string // route path
	statusCode      int    // from status_code=
	responseModel   string // from response_model=
	responseClass   string // from response_class= (e.g. "HTMLResponse")
	includeInSchema bool   // false only when include_in_schema=False literal; default true
}
