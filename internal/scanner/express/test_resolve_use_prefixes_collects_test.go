//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveUsePrefixes_Collects 테스트
package express

import "testing"

func TestResolveUsePrefixes_Collects(t *testing.T) {
	fi := mustParse(t, []byte(`app.use('/api', userRouter); doStuff();`))
	mounts := resolveUsePrefixes(fi, map[string]bool{"app": true}, map[string]string{"userRouter": "./u.ts"}, "", nil)
	if len(mounts) != 1 || mounts[0].Prefix != "/api" || mounts[0].VarName != "userRouter" {
		t.Fatalf("got %+v", mounts)
	}
}
