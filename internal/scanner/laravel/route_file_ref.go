//ff:type feature=scan type=model topic=laravel
//ff:what Provider가 로드하는 분할 라우트 파일 참조(상대경로+초기 prefix/middleware)
package laravel

// routeFileRef references a split route file loaded by a RouteServiceProvider,
// together with the prefix/middleware the provider applies to it.
type routeFileRef struct {
	relPath    string
	prefix     string
	middleware []string
}
