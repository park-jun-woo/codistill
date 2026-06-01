//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractMethodPath_None 테스트
package quarkus

import "testing"

func TestExtractMethodPath_None(t *testing.T) {
	root, src := parseQ(t, `class R { @GET public String get() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if got := extractMethodPath(m, &fileInfo{src: src}); got != "" {
		t.Fatalf("got %q", got)
	}
}
