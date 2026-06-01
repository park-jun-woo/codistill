//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractBinaryLiteralPrefix 테스트
package nestjs

import "testing"

func TestExtractBinaryLiteralPrefix(t *testing.T) {
	src := []byte(`app.setGlobalPrefix(urlPrefix + 'api');`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findAllByType(root, "call_expression")[0]
	args := findChildByType(call, "arguments")
	prefix, ok := extractBinaryLiteralPrefix(args, root, src)
	if !ok || prefix != "api" {
		t.Fatalf("expected api, got %q ok=%v", prefix, ok)
	}
}
