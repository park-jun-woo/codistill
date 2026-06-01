//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectURIVersioning_NoMain 테스트
package nestjs

import "testing"

func TestDetectURIVersioning_NoMain(t *testing.T) {
	if enabled, _ := detectURIVersioning(t.TempDir()); enabled {
		t.Fatal("expected false when no main.ts")
	}
}
