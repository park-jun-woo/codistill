//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveSetPrefixIdentifier 테스트
package nestjs

import "testing"

func TestResolveSetPrefixIdentifier(t *testing.T) {
	src := []byte(`const globalPrefix = 'api'; app.setGlobalPrefix(globalPrefix);`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	var call = findAllByType(root, "call_expression")[0]
	args := findChildByType(call, "arguments")
	prefix, ok := resolveSetPrefixIdentifier(args, root, src)
	if !ok || prefix != "api" {
		t.Fatalf("expected api, got %q ok=%v", prefix, ok)
	}
}
