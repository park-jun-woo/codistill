//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractUseMount_Valid 테스트
package express

import "testing"

func TestExtractUseMount_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter);`))
	m := extractUseMount(firstCallExpr(t, fi), fi, map[string]bool{"app": true}, map[string]string{"userRouter": "./user.ts"}, "", nil)
	if m == nil || m.Prefix != "/api" || m.VarName != "userRouter" || m.SourceRouter != "app" || m.FilePath != "./user.ts" {
		t.Fatalf("got %+v", m)
	}
}
