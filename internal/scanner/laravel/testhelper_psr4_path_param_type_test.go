//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what psr4PathParamType 테스트 헬퍼: 엔드포인트 path param의 타입을 반환한다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

func psr4PathParamType(ep *scanner.Endpoint, name string) string {
	if ep == nil || ep.Request == nil {
		return ""
	}
	for _, p := range ep.Request.PathParams {
		if p.Name == name {
			return p.Type
		}
	}
	return ""
}
