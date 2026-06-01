//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestDedupePaths 테스트
package nestjs

import "testing"

func TestDedupePaths(t *testing.T) {
	in := []string{"a", "b", "a", "c", "b"}
	out := dedupePaths(in)
	want := []string{"a", "b", "c"}
	if len(out) != len(want) {
		t.Fatalf("expected %v, got %v", want, out)
	}
	for i := range want {
		if out[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, out)
		}
	}
}
