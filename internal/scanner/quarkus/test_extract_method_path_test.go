//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractMethodPath 테스트
package quarkus

import "testing"

func TestExtractMethodPath(t *testing.T) {
	root, src := parseQ(t, `class R { @GET @Path("/{id}") public String get() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if got := extractMethodPath(m, &fileInfo{src: src}); got != "/{id}" {
		t.Fatalf("got %q", got)
	}
}
