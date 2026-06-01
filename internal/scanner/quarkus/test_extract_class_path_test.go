//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractClassPath 테스트
package quarkus

import "testing"

func TestExtractClassPath(t *testing.T) {
	root, src := parseQ(t, `@Path("/users") class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPath(cls, &fileInfo{src: src}); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
