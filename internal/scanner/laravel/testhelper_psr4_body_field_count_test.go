//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what psr4BodyFieldCount 테스트 헬퍼: 엔드포인트 요청 본문 필드 수를 반환한다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

func psr4BodyFieldCount(ep *scanner.Endpoint) int {
	if ep == nil || ep.Request == nil || ep.Request.Body == nil {
		return 0
	}
	return len(ep.Request.Body.Fields)
}
