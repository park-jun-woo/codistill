//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestComposerPSR4_Missing 테스트
package laravel

import "testing"

func TestComposerPSR4_Missing(t *testing.T) {
	if m := composerPSR4(t.TempDir()); len(m) != 0 {
		t.Errorf("expected empty map for missing composer.json, got %v", m)
	}
}
