//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveUsePrefixes_None 테스트
package express

import "testing"

func TestResolveUsePrefixes_None(t *testing.T) {
	fi := mustParse(t, []byte(`doStuff();`))
	mounts := resolveUsePrefixes(fi, map[string]bool{"app": true}, nil, "", nil)
	if len(mounts) != 0 {
		t.Fatalf("expected none, got %+v", mounts)
	}
}
