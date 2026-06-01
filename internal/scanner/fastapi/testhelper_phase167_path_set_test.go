//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what Phase167: 스캔 결과의 endpoint path 집합을 만드는 테스트 헬퍼
package fastapi

import "testing"

func phase167PathSet(t *testing.T, scanRoot string) map[string]bool {
	t.Helper()
	res, err := Scan(scanRoot)
	if err != nil {
		t.Fatal(err)
	}
	set := map[string]bool{}
	for _, ep := range res.Endpoints {
		set[ep.Path] = true
	}
	return set
}
