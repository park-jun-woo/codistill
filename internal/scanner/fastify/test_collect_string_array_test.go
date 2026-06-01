//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectStringArray 테스트
package fastify

import "testing"

func TestCollectStringArray(t *testing.T) {
	fi := mustParse(t, []byte(`const x = ["a", "b", c];`+"\n"))
	arrs := findAllByType(fi.Root, "array")
	if len(arrs) == 0 {
		t.Fatal("no array")
	}
	got := collectStringArray(arrs[0], fi.Src)
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("expected [a b] (idents skipped), got %v", got)
	}
}
