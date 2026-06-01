//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractUseMount_NoIdentifierObject 테스트
package express

import "testing"

func TestExtractUseMount_NoIdentifierObject(t *testing.T) {

	fi := mustParse(t, []byte(`a.b.use('/x', r);`))
	if m := extractUseMount(firstCallExpr(t, fi), fi, map[string]bool{"a": true, "b": true}, nil, "", nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
