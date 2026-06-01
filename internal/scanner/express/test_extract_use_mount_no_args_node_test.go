//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractUseMount_NoArgsNode 테스트
package express

import "testing"

func TestExtractUseMount_NoArgsNode(t *testing.T) {
	fi := mustParse(t, []byte("app.use`x`;"))
	if m := extractUseMount(firstCallExpr(t, fi), fi, map[string]bool{"app": true}, nil, "", nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
