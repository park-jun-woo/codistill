//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestTrySetGlobalPrefix_NoMatch 테스트
package nestjs

import "testing"

func TestTrySetGlobalPrefix_NoMatch(t *testing.T) {
	src := []byte(`app.listen(3000);`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call_expression")
	for _, call := range calls {
		_, ok := trySetGlobalPrefix(call, root, src)
		if ok {
			t.Fatal("expected false")
		}
	}
}
