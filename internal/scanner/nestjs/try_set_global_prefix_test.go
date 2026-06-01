//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestTrySetGlobalPrefix_Match 테스트
package nestjs

import "testing"

func TestTrySetGlobalPrefix_Match(t *testing.T) {
	src := []byte(`app.setGlobalPrefix('api');`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no calls found")
	}
	prefix, ok := trySetGlobalPrefix(calls[0], root, src)
	if !ok || prefix != "api" {
		t.Fatalf("expected api, got %q ok=%v", prefix, ok)
	}
}
