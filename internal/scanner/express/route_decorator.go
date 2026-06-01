//ff:type feature=scan type=model topic=express
//ff:what 파싱된 라우트 데코레이터의 이름과 첫 문자열 인자(path)를 담는다
package express

// routeDecorator holds a parsed class/method decorator name and its first
// string argument (used as the route path). hasArg distinguishes @Get('/x')
// from a bare @Get.
type routeDecorator struct {
	name   string
	arg    string
	hasArg bool
}
