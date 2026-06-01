//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestParseUseMountArgs_PrefixNotString 테스트
package express

import "testing"

func TestParseUseMountArgs_PrefixNotString(t *testing.T) {
	fi := mustParse(t, []byte(`app.use(mw, userRouter);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if m := parseUseMountArgs(args, fi, nil, "", nil); m != nil {
		t.Fatalf("got %+v", m)
	}
}
