//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryParseRouterAssignment_NonRouter 테스트
package fastapi

import "testing"

func TestTryParseRouterAssignment_NonRouter(t *testing.T) {

	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	for _, a := range findAllByType(root, "assignment") {
		if ri := tryParseRouterAssignment(a, src, nil); ri != nil {
			t.Fatalf("non-call assignment should be nil, got %v", ri)
		}
	}
}
