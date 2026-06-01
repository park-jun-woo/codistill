//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveConstString 테스트
package express

import "testing"

func TestResolveConstString(t *testing.T) {
	fi := mustParse(t, []byte(`const BASE_API_PATH = '/ghost/api';`))
	if got := resolveConstString(fi.Root, fi.Src, "BASE_API_PATH"); got != "/ghost/api" {
		t.Fatalf("got %q", got)
	}
	if got := resolveConstString(fi.Root, fi.Src, "MISSING"); got != "" {
		t.Fatalf("expected empty for missing, got %q", got)
	}
}
