//ff:type feature=scan type=model topic=fastapi
//ff:what include_router 호출 정보 구조체
package fastapi

// includeCall represents an app.include_router(router, prefix=...) call.
type includeCall struct {
	parentVar   string // e.g., "app"
	childVar    string // e.g., "router" (attribute이면 변수 부분만)
	childModule string // e.g., "items" (attribute이면 모듈 부분, 단순 identifier이면 "")
	extraPrefix string // prefix kwarg from include_router call
	hidden      bool   // true if include_router(..., include_in_schema=False)
}
