//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFirstPSR4Dir 테스트
package laravel

import (
	"encoding/json"
	"testing"
)

func TestFirstPSR4Dir(t *testing.T) {
	if got := firstPSR4Dir(json.RawMessage(`"app/"`)); got != "app/" {
		t.Errorf("string form = %q, want app/", got)
	}
	if got := firstPSR4Dir(json.RawMessage(`["src/", "lib/"]`)); got != "src/" {
		t.Errorf("array form = %q, want src/", got)
	}
	if got := firstPSR4Dir(json.RawMessage(`[]`)); got != "" {
		t.Errorf("empty array = %q, want empty", got)
	}
	if got := firstPSR4Dir(json.RawMessage(`123`)); got != "" {
		t.Errorf("number = %q, want empty", got)
	}
}
