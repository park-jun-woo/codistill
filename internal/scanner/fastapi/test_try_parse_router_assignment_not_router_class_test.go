//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryParseRouterAssignment_NotRouterClass 테스트
package fastapi

import "testing"

func TestTryParseRouterAssignment_NotRouterClass(t *testing.T) {

	src := []byte("db = SomeClass()\n")
	root, _ := parsePython(src)
	for _, a := range findAllByType(root, "assignment") {
		if ri := tryParseRouterAssignment(a, src, nil); ri != nil {
			t.Fatalf("non-router class should be nil, got %v", ri)
		}
	}
}
