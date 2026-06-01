//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestParseUseMountArgs_Valid 테스트
package express

import "testing"

func TestParseUseMountArgs_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	m := parseUseMountArgs(args, fi, map[string]string{"userRouter": "./u.ts"}, "", nil)
	if m == nil || m.Prefix != "/api" || m.VarName != "userRouter" || m.FilePath != "./u.ts" {
		t.Fatalf("got %+v", m)
	}
}
