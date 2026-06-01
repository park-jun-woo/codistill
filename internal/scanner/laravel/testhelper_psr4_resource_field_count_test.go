//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what psr4ResourceFieldCount 테스트 헬퍼: 지정 TypeName 응답의 필드 수를 반환한다
package laravel

import "github.com/park-jun-woo/codistill/internal/scanner"

func psr4ResourceFieldCount(ep *scanner.Endpoint, typeName string) int {
	if ep == nil {
		return 0
	}
	for _, r := range ep.Responses {
		if r.TypeName == typeName {
			return len(r.Fields)
		}
	}
	return 0
}
